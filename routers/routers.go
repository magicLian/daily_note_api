package routers

import (
	log "github.com/sirupsen/logrus"
	"ml_daily_record/pkg/configs"
	"ml_daily_record/pkg/controllers"
	"ml_daily_record/pkg/middleware"

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

	auth := v1.Group("/auth")
	{
		authC := new(controllers.AuthC)
		auth.POST("/native", authC.AuthNative)
	}

	dn := v1.Group("/dailyNotes", middleware.TokenAuth)
	{
		dnC := new(controllers.DailyNoteC)
		dn.GET("", dnC.GetDailyNotes)
		dn.POST("", dnC.CreateDailyNotes)
		dn.PUT("", dnC.UpdateDailyNotes)
		dn.DELETE("/:id", dnC.DeleteDailyNotes)
	}

	return router
}
