package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"ml_daily_record/pkg/configs"
	"ml_daily_record/pkg/db"
	"ml_daily_record/pkg/models"
	"ml_daily_record/pkg/services"
	"ml_daily_record/pkg/util"
	"net/http"
	"strings"
)

func TokenAuth(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if auth != "" {
		if !strings.HasPrefix(auth, "Bearer ") {
			util.HttpResult(c, http.StatusUnauthorized, models.AuthTokenNotFound, nil)
			return
		} else {
			auth = strings.Split(auth, "Bearer ")[1]
		}
	} else {
		cookie, err := c.Request.Cookie("EIToken")
		if err != nil {
			util.HttpResult(c, http.StatusUnauthorized, models.AuthTokenNotFound, nil)
			return
		}
		auth = cookie.Value
	}

	log.Debugf("token is [%s]", auth)
	if auth == "" {
		util.HttpResult(c, http.StatusUnauthorized, models.AuthTokenNotFound, nil)
		return
	}

	authMap, err := services.ParseToken(auth, configs.Cf.Jwt.Secret)
	if err != nil {
		util.HttpResult(c, http.StatusUnauthorized, models.AuthTokenNotFound, nil)
		return
	}
	b, err := services.IsTokenExpired(authMap["exp"].(string))
	if err != nil {
		util.HttpResult(c, http.StatusUnauthorized, err, nil)
		return
	}
	if !b {
		util.HttpResult(c, http.StatusUnauthorized, models.AuthTokenExpired, nil)
		return
	}

	userMap := authMap["user"]
	userBytes, err := json.Marshal(userMap)
	if err != nil {
		util.HttpResult(c, http.StatusUnauthorized, models.UserNotFound, nil)
		return
	}
	tokenUser := &models.TokenUser{}
	if err = json.Unmarshal(userBytes, &tokenUser); err != nil {
		util.HttpResult(c, http.StatusUnauthorized, models.UserNotFound, nil)
		return
	}

	userQuery, err := db.PGDB.FindUserByUsernameAndToken(tokenUser.Username, auth)
	if err != nil {
		util.HttpResult(c, http.StatusUnauthorized, models.UserNotFound, nil)
		return
	}
	if userQuery == nil {
		util.HttpResult(c, http.StatusUnauthorized, models.AuthTokenInvaild, nil)
		return
	}
	if userQuery.Token == "" || userQuery.Token != auth {
		util.HttpResult(c, http.StatusUnauthorized, models.AuthTokenInvaild, nil)
		return
	}

	c.Set("loginUser", userQuery)
	c.Next()
}
