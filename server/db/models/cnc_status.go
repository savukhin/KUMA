package models

import (
	"errors"
	"strings"
)

type CncStatus struct {
	ID         uint64 `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	StatusName string `gorm:"column:status_name;type:varchar(70);unique" json:"status_name"`
}

var (
	WorkingStatus = &CncStatus{ID: 0, StatusName: "Working"}
	StoppedStatus = &CncStatus{ID: 1, StatusName: "Stopped"}
	BrokenStatus  = &CncStatus{ID: 2, StatusName: "Broken"}
)

func StringToStatusID(s string) (uint64, error) {
	lower := strings.ToLower(s)

	switch lower {
	case "working":
		return 0, nil
	case "stopped":
		return 1, nil
	case "broken":
		return 2, nil
	}

	return 0, errors.New("no such status " + s)
}
