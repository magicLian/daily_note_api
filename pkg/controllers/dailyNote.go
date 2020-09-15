package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	_const "ml_daily_record/pkg/const"
	"ml_daily_record/pkg/models"
	"ml_daily_record/pkg/services"
	"ml_daily_record/pkg/util"
	"net/http"
)

type DailyNoteC struct {
}

func (dnC *DailyNoteC) GetDailyNotes(c *gin.Context) {
	dnQueryReq := &models.DailyNoteQueryReq{}

	from := c.Query("from")
	to := c.Query("to")
	if (from == "" && to != "") || (from != "" && to == "") {
		util.HttpResult(c, http.StatusBadRequest, models.DailyQueryParametersInvaild, nil)
		return
	}
	dnQueryReq.From = from
	dnQueryReq.To = to

	t := c.Query("type")
	if t != "" {
		if t != _const.DAILY_TYPE_ANNIVERSARY && t != _const.DAILY_TYPE_BIRTHDAY &&
			t != _const.DAILY_TYPE_HOILDAY && t != _const.DAILY_TYPE_MANUAL_SIGNED &&
			t != _const.DAILY_TYPE_PEROID {
			log.Errorf("type invaild:[%s]", t)
			util.HttpResult(c, http.StatusBadRequest, models.DailyQueryParametersInvaild, nil)
			return
		}
		dnQueryReq.Type = t
	}

	dnList, err := services.GetDailyNotes(dnQueryReq)
	util.Response(c, err, dnList, int64(len(dnList)))
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

func (dnC *DailyNoteC) UpdateDailyNote(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		util.HttpResult(c, http.StatusBadRequest, models.DailyNoteIdNotFound, nil)
		return
	}

	dnNew := &models.DailyNote{}
	if err := c.BindJSON(&dnNew); err != nil {
		log.Errorf(err.Error())
		util.HttpResult(c, http.StatusBadRequest, models.DailyNoteParseFailed, nil)
		return
	}
	if dnNew.Date == "" {
		util.HttpResult(c, http.StatusBadRequest, models.DailyNoteParseFailed, nil)
		return
	}
	if len(dnNew.Type) == 0 {
		util.HttpResult(c, http.StatusBadRequest, models.DailyNoteParseFailed, nil)
		return
	}
	if dnNew.Level == "" {
		util.HttpResult(c, http.StatusBadRequest, models.DailyNoteParseFailed, nil)
		return
	}
	err := services.UpdateDailyNote(dnNew)
	util.Response(c, err, dnNew, 0)
}

func (dnC *DailyNoteC) DeleteDailyNote(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		util.HttpResult(c, http.StatusBadRequest, models.DailyNoteIdNotFound, nil)
		return
	}
	err := services.DeleteDailyNoteById(id)
	util.Response(c, err, nil, 0)
}
