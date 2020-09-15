package services

import (
	uuid "github.com/satori/go.uuid"
	"ml_daily_record/pkg/db"
	"ml_daily_record/pkg/models"
)

func GetDailyNotes(dnQueryReq *models.DailyNoteQueryReq) ([]*models.DailyNote, error) {
	return db.PGDB.FindDailyNoteByFilter(dnQueryReq)
}

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

func UpdateDailyNote(dnNew *models.DailyNote) error {
	dnTarget, err := db.PGDB.FindDailyNoteById(dnNew.ID)
	if err != nil {
		return err
	}
	if dnNew.Date != dnTarget.Date {
		return models.DailyNoteUpdateConfict
	}
	if err := db.PGDB.UpdateDailyNote(dnNew); err != nil {
		return err
	}

	return nil
}

func DeleteDailyNoteById(id string) error {
	dnQuery, err := db.PGDB.FindDailyNoteById(id)
	if err != nil {
		return err
	}
	if dnQuery == nil {
		return models.DailyNoteNotFound
	}
	if err := db.PGDB.DeleteDailyNote(id); err != nil {
		return err
	}

	return nil
}
