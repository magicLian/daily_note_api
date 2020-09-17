package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	_const "ml_daily_record/pkg/const"
	"ml_daily_record/pkg/models"
	"ml_daily_record/pkg/services"
	"ml_daily_record/pkg/util"
	"net/http"
	"strings"
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

func (dnC *DailyNoteC) CreateDailyNotes(c *gin.Context) {
	dnList := make([]*models.DailyNote, 0)
	if err := c.BindJSON(&dnList); err != nil {
		log.Errorf(err.Error())
		util.HttpResult(c, http.StatusBadRequest, models.DailyNoteParseFailed, nil)
		return
	}
	for _, dn := range dnList {
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
	}

	err := services.CreateDailyNotes(dnList)
	util.Response(c, err, dnList, int64(len(dnList)))
}

func (dnC *DailyNoteC) UpdateDailyNotes(c *gin.Context) {
	dnNewList := make([]*models.DailyNote, 0)
	if err := c.BindJSON(&dnNewList); err != nil {
		log.Errorf(err.Error())
		util.HttpResult(c, http.StatusBadRequest, models.DailyNoteParseFailed, nil)
		return
	}
	for _, dnNew := range dnNewList {
		if dnNew.ID == "" {
			util.HttpResult(c, http.StatusBadRequest, models.DailyNoteIdNotFound, nil)
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
	}

	err := services.UpdateDailyNotes(dnNewList)
	util.Response(c, err, dnNewList, int64(len(dnNewList)))
}

func (dnC *DailyNoteC) DeleteDailyNotes(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		util.HttpResult(c, http.StatusBadRequest, models.DailyNoteIdNotFound, nil)
		return
	}

	idArr := strings.Split(idStr, ",")
	if len(idStr) == 0 {
		util.HttpResult(c, http.StatusBadRequest, models.DailyNoteIdNotFound, nil)
		return
	}

	err := services.DeleteDailyNoteByIds(idArr)
	util.Response(c, err, nil, 0)
}
