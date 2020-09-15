package models

import (
	"errors"
	"github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

var (
	DailyNoteExist              = errors.New("daily note already exist")
	DailyNoteParseFailed        = errors.New("daily note parse failed")
	DailyNoteIdNotFound         = errors.New("daily note id not found")
	DailyNoteNotFound           = errors.New("daily note not found")
	DailyNoteUpdateConfict      = errors.New("daily note update confict")
	DailyQueryParametersInvaild = errors.New("daily note query parameters invaild")
)

type DailyNote struct {
	ID        string         `json:"id" gorm:"type:varchar(50);primary_key"`
	Title     string         `json:"title" gorm:"type:varchar(100)"`
	Date      string         `json:"date" gorm:"type:varchar(50)"`
	Note      string         `json:"note" gorm:"type:varchar(255)"`
	Level     string         `json:"level" gorm:"type:varchar(5)"`
	Type      []string       `json:"type" gorm:"-"`
	TypeStr   string         `json:"-"`
	RawData   postgres.Jsonb `json:"rawData"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt *time.Time     `json:"deletedAt"`
}

type PeriodData struct {
	IsStart bool `json:"isStart"`
	IsEnd   bool `json:"isEnd"`
}

type DailyNoteQueryReq struct {
	From string `json:"from"`
	To   string `json:"to"`
	Type string `json:"type"`
}
