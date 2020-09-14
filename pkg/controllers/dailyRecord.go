package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"ml_daily_record/pkg/models"
	"ml_daily_record/pkg/services"
	"ml_daily_record/pkg/util"
	"net/http"
)

type DailyNoteC struct {
}

func (dnC *DailyNoteC) CreateDailyNote(c *gin.Context) {
	dn := &models.DailyNote{}
	if err := c.BindJSON(&dn); err != nil {
		log.Errorf(err.Error())
		util.HttpResult(c, http.StatusBadRequest, models.DailyNoteParseFailed, nil)
		return
	}
	if dn.Date == "" {
		util.HttpResult(c, http.StatusBadRequest, models.DailyNoteParseFailed, nil)
		return
	}
	if len(dn.Type) == 0 {
		util.HttpResult(c, http.StatusBadRequest, models.DailyNoteParseFailed, nil)
		return
	}
	if dn.Level == "" {
		util.HttpResult(c, http.StatusBadRequest, models.DailyNoteParseFailed, nil)
		return
	}
	err := services.CreateDailyNote(dn)
	util.Response(c, err, dn, 0)
}
