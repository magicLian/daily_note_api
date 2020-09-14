package util

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

func HttpResult(c *gin.Context, code int, err error, data interface{}) {
	if err != nil {
		log.Error(err.Error())
		c.JSON(code, gin.H{
			"timestamp": time.Now().Format("2006-01-02T15:04:05Z07:00"),
			"status":    code,
			"error":     err.Error(),
			"data":      data,
			"path":      c.Request.RequestURI,
		})
	} else {
		c.JSON(code, gin.H{
			"timestamp": time.Now().Format("2006-01-02T15:04:05Z07:00"),
			"status":    code,
			"error":     nil,
			"data":      data,
			"path":      c.Request.RequestURI,
		})
	}
	c.Abort()
}

func HttpResultWithTotalCounts(c *gin.Context, code int, err error, data interface{}, counts int64) {
	if err != nil {
		log.Error(err.Error())
		c.JSON(code, gin.H{
			"timestamp": time.Now().Format("2006-01-02T15:04:05Z07:00"),
			"status":    code,
			"error":     err.Error(),
			"data":      data,
			"path":      c.Request.RequestURI,
		})
	} else {
		c.JSON(code, gin.H{
			"timestamp":  time.Now().Format("2006-01-02T15:04:05Z07:00"),
			"status":     code,
			"error":      nil,
			"data":       data,
			"totalCount": counts,
			"path":       c.Request.RequestURI,
		})
	}
	c.Abort()
}

func Response(c *gin.Context, e error, d interface{}, count int64) {
	if e != nil {
		if strings.Contains(e.Error(), "permission") {
			HttpResult(c, 403, e, nil)
			return
		}
		if strings.Contains(e.Error(), "not exist") {
			HttpResult(c, 404, e, nil)
			return
		}
		if strings.Contains(e.Error(), "already exist") {
			HttpResult(c, 409, e, nil)
			return
		}
		if strings.Contains(e.Error(), "forbidden") {
			HttpResult(c, 403, e, nil)
			return
		}
		HttpResult(c, 500, e, nil)
		return
	}
	if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPatch {
		HttpResult(c, 201, nil, d)
	} else if c.Request.Method == http.MethodDelete {
		HttpResult(c, 204, nil, d)
	} else {
		HttpResultWithTotalCounts(c, 200, nil, d, count)
	}
	return
}
