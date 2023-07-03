package mock_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"os"
	"time"

	"path/filepath"
	"server/api"
	"server/db"
	"server/db/models"
	"server/db/query"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetExecutablePath() string {
	ex, _ := os.Executable()

	return filepath.Dir(ex)
}

func TestFull(t *testing.T) {
	gormDB, err := gorm.Open(sqlite.Open(GetExecutablePath()+"/test.db"), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 	"localhost", "20624880", "admin", "kuma", "5432",
	// )
	// postgresConn := postgres.Open(dsn)
	// gormDB, err := gorm.Open(postgresConn)

	require.NoError(t, err)

	err = gormDB.AutoMigrate(models.MigrateModels...)
	require.NoError(t, err)
	// gormDB.CreateTab

	db.InitCncStatusTable(gormDB)

	logger, err := zap.NewDevelopment()
	// logger, err := zap.Config{}.Build()
	require.NoError(t, err)

	stopped := make([]string, 0)
	app := api.SetupRouter(gormDB, []byte("test-key"), logger, func(s string) { stopped = append(stopped, s) })

	req := httptest.NewRequest("GET", "/ping", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "Pong", string(body))

	// ------- CREATE CNC CHECKER ------- //

	cncCheckerUsername := "cool_cnc"
	cncCheckerPassword := "cool_cnc"
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(cncCheckerPassword), bcrypt.DefaultCost)
	require.NoError(t, err)

	cncChecker := &models.CncChecker{
		// ID: 0,
		Credentials: models.Credentials{
			Username:     cncCheckerUsername,
			PasswordHash: string(passwordHash),
		},
		Title:    "main",
		StatusID: models.StoppedStatus.ID,
	}

	passwordHash, _ = bcrypt.GenerateFromPassword([]byte(cncCheckerPassword+"2"), bcrypt.DefaultCost)
	cncChecker2 := &models.CncChecker{
		// ID: 0,
		Credentials: models.Credentials{
			Username:     cncCheckerUsername + "2",
			PasswordHash: string(passwordHash),
		},
		Title:    "main2",
		StatusID: models.StoppedStatus.ID,
	}
	cnc := query.Use(gormDB).CncChecker
	cnt, err := cnc.Count()
	require.NoError(t, err)
	require.EqualValues(t, 0, cnt)

	err = cnc.Create(cncChecker)
	require.NoError(t, err)
	err = cnc.Create(cncChecker2)
	require.NoError(t, err)

	cnt, err = cnc.Count()
	require.NoError(t, err)
	require.EqualValues(t, 2, cnt)

	// ------- LOGIN CNC CHECKER ------- //

	payload := &api.LoginUserDTO{
		Username: cncCheckerUsername,
		Password: cncCheckerPassword,
	}

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(payload)
	require.NoError(t, err)

	req = httptest.NewRequest("GET", "/api/v1/auth/login", &buf)
	req.Header.Set("Content-Type", fiber.MIMEApplicationJSON)

	resp, err = app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)

	body, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	tokens := make(map[string]string)
	json.Unmarshal(body, &tokens)

	accessToken, ok := tokens[api.AccessTokenHeaderName]
	require.True(t, ok)
	_, ok = tokens[api.RefreshTokenHeaderName]
	require.True(t, ok)

	req = httptest.NewRequest("GET", "/api/v1/auth/restricted", nil)
	req.Header.Set("Content-Type", fiber.MIMEApplicationJSON)
	req.Header.Set(api.AuthorizationHeaderName, "Bearer "+accessToken)

	resp, err = app.Test(req)
	require.NoError(t, err)

	body, err = io.ReadAll(resp.Body)
	require.NoError(t, err)

	require.Equal(t, fiber.StatusOK, resp.StatusCode)
	require.Equal(t, "Welcome 1", string(body))

	// ------- STOP / WORK CNC CHECKER ------- //

	payloadUpdate := &api.UpdateCncDTO{
		Status: "stopped",
	}

	err = json.NewEncoder(&buf).Encode(payloadUpdate)
	require.NoError(t, err)

	require.Len(t, stopped, 0)

	req = httptest.NewRequest("POST", "/api/v1/cnc/update-status", &buf)
	req.Header.Set("Content-Type", fiber.MIMEApplicationJSON)
	req.Header.Set(api.AuthorizationHeaderName, "Bearer "+accessToken)

	resp, err = app.Test(req)
	require.NoError(t, err)

	require.Equal(t, fiber.StatusOK, resp.StatusCode)

	require.Len(t, stopped, 1)
	require.Equal(t, "main", stopped[0])

}
