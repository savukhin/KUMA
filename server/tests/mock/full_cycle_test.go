package mock_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"os"

	"path/filepath"
	"server/api"
	"server/db/models"
	"server/db/query"
	"testing"
	"time"

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
	db, err := gorm.Open(sqlite.Open(GetExecutablePath()+"/test.db"), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	require.NoError(t, err)

	err = db.AutoMigrate(models.MigrateModels...)
	require.NoError(t, err)

	logger, err := zap.NewDevelopment()
	// logger, err := zap.Config{}.Build()
	require.NoError(t, err)

	app := api.SetupRouter(db, []byte("test-key"), logger, nil)

	req := httptest.NewRequest("GET", "/ping", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "Pong", string(body))

	cncCheckerUsername := "cool_cnc"
	cncCheckerPassword := "cool_cnc"
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(cncCheckerPassword), bcrypt.DefaultCost)
	require.NoError(t, err)

	cncChecker := &models.CncChecker{
		Credentials: models.Credentials{
			Username:     cncCheckerUsername,
			PasswordHash: string(passwordHash),
		},
	}
	cnc := query.Use(db).CncChecker
	cnt, err := cnc.Count()
	require.NoError(t, err)
	require.EqualValues(t, 0, cnt)

	err = cnc.Create(cncChecker)
	require.NoError(t, err)

	payload := &api.LoginUserDTO{
		Username: cncCheckerUsername,
		Password: cncCheckerPassword,
	}

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(payload)
	require.NoError(t, err)

	// req = httptest.NewRequest("GET", "/api/v1/auth/login", strings.NewReader(`{"username":"cool_cnc","password":"cool_cnc"}`))
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
	require.Equal(t, "Welcome 0", string(body))

}
