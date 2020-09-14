package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"ml_daily_record/pkg/configs"
)

var x *gorm.DB

func InitPg() {
	log.Info("starting init db")
	pgInfo := "user=" + configs.Cf.Db.Postgres.Username +
		" password=" + configs.Cf.Db.Postgres.Password +
		" dbname=" + configs.Cf.Db.Postgres.DBName +
		" host=" + configs.Cf.Db.Postgres.Host +
		" port=" + configs.Cf.Db.Postgres.Port +
		" sslmode=disable"

	log.Infof("pgInfo : %s", pgInfo)

	err := errors.New("")
	x, err = gorm.Open("postgres", pgInfo)
	if err != nil {
		log.Fatalf("Fail to connect pg: %v\n", err)
	}

	x.LogMode(false)
	x.SingularTable(true)
	x.DB().SetMaxIdleConns(configs.Cf.Db.Postgres.MinPoolSize)
	x.DB().SetMaxOpenConns(configs.Cf.Db.Postgres.MaxPoolSize)

	initTable()
	initData()
}

func initTable() {
	//x.AutoMigrate(&models.ChartRepo{})
}

func initData() {
}

func GetDbInstance() *gorm.DB {
	return x
}
