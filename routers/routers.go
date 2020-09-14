package routers

import (
	log "github.com/sirupsen/logrus"
	"ml_daily_record/pkg/configs"
	"ml_daily_record/pkg/controllers"

	//"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	log.Info("starting init routers")

	if configs.Cf.Deploy != "dev" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	router := gin.New()
	router.Use(gin.Recovery())
	gin.DisableConsoleColor()
	//pprof.Register(router)

	router.Static("/apidoc", "./resources/apidoc")

	v1 := router.Group("/v1")

	dn := v1.Group("/dailyNotes")
	{
		dnC := new(controllers.DailyNoteC)
		dn.POST("", dnC.CreateDailyNote)
	}

	return router
}
