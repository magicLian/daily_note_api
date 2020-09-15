package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"ml_daily_record/pkg/models"
	"ml_daily_record/pkg/services"
	"ml_daily_record/pkg/util"
	"net/http"
)

type AuthC struct {
}

func (ac *AuthC) AuthNative(c *gin.Context) {
	an := &models.AuthNative{}
	if err := c.BindJSON(&an); err != nil {
		log.Errorf(err.Error())
		util.HttpResult(c, http.StatusBadRequest, models.AuthNativeParseFailed, nil)
		return
	}

	if an.Username == "" {
		util.HttpResult(c, http.StatusBadRequest, models.AuthNativeUsernameNotFound, nil)
		return
	}
	if an.Password == "" {
		util.HttpResult(c, http.StatusBadRequest, models.AuthNativePasswordNotFound, nil)
		return
	}

	token, err := services.AuthNative(an)
	util.Response(c, err, token, 0)
}
