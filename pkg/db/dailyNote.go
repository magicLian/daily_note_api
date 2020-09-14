package db

import (
	"encoding/json"
	"fmt"
	"ml_daily_record/pkg/models"
	"time"
)

type DailyNoteDao interface {
	CreateDailyNote(*models.DailyNote) error
	BatchCreateDailyNotes([]*models.DailyNote) error
	UpdateDailyNote(*models.DailyNote) error
	BatchUpdateDailyNotes([]*models.DailyNote) error
	DeleteDailyNote(string) error
	FindDailyNoteById(string) (*models.DailyNote, error)
	FindDailyNoteByDateDuration(time.Time, time.Time) ([]*models.DailyNote, error)
	ExistDailyNote(*models.DailyNote) (bool, error)
}

func (pg *PGService) CreateDailyNote(dn *models.DailyNote) error {
	if err := changeDailyNoteToString(dn); err != nil {
		return err
	}
	if err := pg.Connection.Create(&dn).Error; err != nil {
		return err
	}
	return nil
}

func (pg *PGService) BatchCreateDailyNotes(dns []*models.DailyNote) error {
	for _, dn := range dns {
		if err := pg.CreateDailyNote(dn); err != nil {
			return err
		}
	}

	return nil
}

func (pg *PGService) UpdateDailyNote(dn *models.DailyNote) error {
	if err := changeDailyNoteToString(dn); err != nil {
		return err
	}
	if err := pg.Connection.Update(&dn).Error; err != nil {
		return err
	}

	return nil
}

func (pg *PGService) BatchUpdateDailyNotes(dns []*models.DailyNote) error {
	for _, dn := range dns {
		if err := pg.UpdateDailyNote(dn); err != nil {
			return err
		}
	}

	return nil
}

func (pg *PGService) DeleteDailyNote(id string) error {
	if err := pg.Connection.Where("id = ?", id).Delete(&models.DailyNote{}).Error; err != nil {
		return err
	}

	return nil
}

func (pg *PGService) FindDailyNoteById(id string) (*models.DailyNote, error) {
	dn := &models.DailyNote{}
	if err := pg.Connection.Where("id = ?", id).First(&dn); err != nil {
		if err.RecordNotFound() {
			return nil, nil
		} else {
			return nil, err.Error
		}
	}
	if err := changeDailyNoteToStruct(dn); err != nil {
		return nil, err
	}

	return dn, nil
}

func (pg *PGService) FindDailyNoteByDateDuration(from time.Time, to time.Time) ([]*models.DailyNote, error) {
	sql := fmt.Sprintf("date between TIMESTAMP '%s' and TIMESTAMP '%s'", from.String(), to.String())
	dailyNotes := make([]*models.DailyNote, 0)

	if err := pg.Connection.Where(sql).Find(&dailyNotes); err != nil {
		if err.RecordNotFound() {
			return nil, nil
		} else {
			return nil, err.Error
		}
	}

	for _, dn := range dailyNotes {
		if err := changeDailyNoteToStruct(dn); err != nil {
			return nil, err
		}
	}

	return dailyNotes, nil
}

func (pg *PGService) ExistDailyNote(dn *models.DailyNote) (bool, error) {
	sql := ""
	if dn.ID != "" {
		sql += fmt.Sprintf("id = '%s'", dn.ID)
	}
	if dn.Date != "" {
		if sql != "" {
			sql += " or "
		}
		sql += fmt.Sprintf("date = '%s'", dn.Date)
	}
	dnQuery := &models.DailyNote{}
	if err := pg.Connection.Where(sql).First(dnQuery); err != nil {
		if err.RecordNotFound() {
			return false, nil
		} else {
			return false, err.Error
		}
	}
	if dnQuery.ID == "" {
		return false, nil
	}

	return true, nil
}

func changeDailyNoteToStruct(dn *models.DailyNote) error {
	if dn.TypeStr != "" {
		typeArr := make([]string, 0)
		if err := json.Unmarshal([]byte(dn.TypeStr), &typeArr); err != nil {
			return err
		}

		dn.Type = typeArr
	}

	return nil
}

func changeDailyNoteToString(dn *models.DailyNote) error {
	typeBytes, err := json.Marshal(dn.Type)
	if err != nil {
		return err
	}

	dn.TypeStr = string(typeBytes)

	return nil
}
