package main

import (
	"ml_daily_record/pkg/configs"
	"ml_daily_record/pkg/db"
	"ml_daily_record/pkg/log"
	"ml_daily_record/routers"
)

func main() {
	//init log
	log.InitLogger()

	//init config
	configs.InitConfig()

	//init db
	db.InitPg()

	//init router
	router := routers.InitRouter()

	//run server
	_ = router.Run(":" + configs.Cf.Server.Port)
}
