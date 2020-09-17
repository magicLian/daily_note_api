package services

import (
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"ml_daily_record/pkg/db"
	"ml_daily_record/pkg/models"
)

func GetDailyNotes(dnQueryReq *models.DailyNoteQueryReq) ([]*models.DailyNote, error) {
	return db.PGDB.FindDailyNoteByFilter(dnQueryReq)
}

func CreateDailyNotes(dnList []*models.DailyNote) error {
	for _, dn := range dnList {
		b, err := db.PGDB.ExistDailyNote(dn)
		if err != nil {
			return err
		}
		if b {
			log.Errorf("create daily note failed, reason: dailyNote already exist, date is:[%s]", dn.Date)
			return models.DailyNoteExist
		}
	}

	tx := db.PGDB.Connection.Begin()
	for _, dn := range dnList {
		dn.ID = uuid.NewV4().String()
		if err := db.PGDB.CreateDailyNoteWithTx(tx, dn); err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()

	return nil
}

func UpdateDailyNotes(dnNewList []*models.DailyNote) error {
	tx := db.PGDB.Connection.Begin()

	for _, dnNew := range dnNewList {
		dnTarget, err := db.PGDB.FindDailyNoteById(dnNew.ID)
		if err != nil {
			return err
		}
		if dnNew.Date != dnTarget.Date {
			return models.DailyNoteUpdateConfict
		}
	}

	if err := db.PGDB.BatchUpdateDailyNotesWithTx(tx, dnNewList); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func DeleteDailyNoteByIds(idArr []string) error {
	for _, id := range idArr {
		dnQuery, err := db.PGDB.FindDailyNoteById(id)
		if err != nil {
			return err
		}
		if dnQuery == nil {
			return models.DailyNoteNotFound
		}
	}

	tx := db.PGDB.Connection.Begin()
	if err := db.PGDB.BatchDeleteDailyNoteByIdWithTx(tx, idArr); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
