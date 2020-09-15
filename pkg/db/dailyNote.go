package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"ml_daily_record/pkg/models"
	"strings"
)

type DailyNoteDao interface {
	CreateDailyNote(*models.DailyNote) error
	BatchCreateDailyNotes([]*models.DailyNote) error
	UpdateDailyNote(*models.DailyNote) error
	BatchUpdateDailyNotes([]*models.DailyNote) error
	DeleteDailyNote(string) error
	FindDailyNoteById(string) (*models.DailyNote, error)
	FindDailyNoteByFilter(*models.DailyNoteQueryReq) ([]*models.DailyNote, error)
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

func (pg *PGService) FindDailyNoteByFilter(filter *models.DailyNoteQueryReq) ([]*models.DailyNote, error) {
	sql := ""

	if filter.From != "" && filter.To != "" {
		sql += fmt.Sprintf("date between '%s' and  '%s'", filter.From, filter.To)
	}
	if filter.Type != "" {
		if sql != "" {
			sql += " and "
		}
		sql += fmt.Sprintf("'%s' = ANY(string_to_array(type_str,','))", filter.Type)
	}

	dailyNotes := make([]*models.DailyNote, 0)

	if err := pg.Connection.Where(sql).Find(&dailyNotes); err.Error != nil {
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
	if err := pg.Connection.Where(sql).First(dnQuery).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}
	}
	if dnQuery.ID == "" {
		return false, nil
	}

	return true, nil
}

func changeDailyNoteToStruct(dn *models.DailyNote) error {
	if dn.TypeStr != "" {
		dn.Type = strings.Split(dn.TypeStr, ",")
	}

	return nil
}

func changeDailyNoteToString(dn *models.DailyNote) error {
	dn.TypeStr = ""
	for i, t := range dn.Type {
		if i == 0 {
			dn.TypeStr += t
		} else {
			dn.TypeStr += "," + t
		}
	}

	return nil
}
