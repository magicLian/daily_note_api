package services

import (
	uuid "github.com/satori/go.uuid"
	"ml_daily_record/pkg/db"
	"ml_daily_record/pkg/models"
)

func CreateDailyNote(dn *models.DailyNote) error {
	b, err := db.PGDB.ExistDailyNote(dn)
	if err != nil {
		return err
	}
	if b {
		return models.DailyNoteExist
	} else {
		dn.ID = uuid.NewV4().String()
		if err := db.PGDB.CreateDailyNote(dn); err != nil {
			return err
		}
	}

	return nil
}
