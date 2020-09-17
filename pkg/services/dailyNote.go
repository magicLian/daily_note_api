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
	mapDate := map[string]int{}

	for _, dn := range dnList {
		b, err := db.PGDB.ExistDailyNote(dn)
		if err != nil {
			return err
		}
		if b {
			log.Errorf("create daily note failed, reason: dailyNote already exist, date is:[%s]", dn.Date)
			return models.DailyNoteExist
		}
		if _, ok := mapDate[dn.Date]; !ok {
			mapDate[dn.Date] = 1
		} else {
			log.Errorf("create daily note failed, reason: duplicated dailyNote")
			return models.DailyNoteDulicated
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
	mapDate := map[string]string{}

	for _, dnNew := range dnNewList {
		if _, ok := mapDate[dnNew.Date]; !ok {
			mapDate[dnNew.Date] = dnNew.ID
		} else {
			log.Errorf("update daily note failed, reason: duplicate dailyNote")
			return models.DailyNoteDulicated
		}

		dnTarget, err := db.PGDB.FindDailyNoteById(dnNew.ID)
		if err != nil {
			return err
		}
		if dnTarget == nil {
			return models.DailyNoteNotFound
		}
		if dnNew.Date != dnTarget.Date {
			return models.DailyNoteUpdateConfict
		}
	}

	tx := db.PGDB.Connection.Begin()
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
