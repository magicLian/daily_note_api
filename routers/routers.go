package routers

import (
	"ml_daily_record/pkg/configs"
	"ml_daily_record/pkg/controllers"
	log "github.com/sirupsen/logrus"
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
	gin.DisableConsoleColor()
	router := gin.Default()
	//pprof.Register(router)

	router.Static("/apidoc", "./resources/apidoc")

	v1 := router.Group("/v1")

	cr := v1.Group("/daily")
	{
		insC := new(controllers.ChartRepoC)
		cr.GET("", insC.GetChartRepos)
	}

	return router
}
