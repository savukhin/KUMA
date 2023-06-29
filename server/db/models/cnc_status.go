package models

import (
	"errors"
	"strings"
)

type CncStatus struct {
	ID         IDType `gorm:"column:id;type:bigint;primaryKey;unique" json:"id"`
	StatusName string `gorm:"column:status_name;type:varchar(70);unique" json:"status_name"`
}

var (
	StoppedStatus = &CncStatus{ID: 1, StatusName: "Stopped"}
	WorkingStatus = &CncStatus{ID: 2, StatusName: "Working"}
	BrokenStatus  = &CncStatus{ID: 3, StatusName: "Broken"}
)

func StringToStatusID(s string) (IDType, error) {
	lower := strings.ToLower(s)

	switch lower {
	case "stopped":
		return 1, nil
	case "working":
		return 2, nil
	case "broken":
		return 3, nil
	}

	return 0, errors.New("no such status " + s)
}
