package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"ml_daily_record/pkg/configs"
	"ml_daily_record/pkg/models"
)

var (
	PGDB PGService
)

type PGService struct {
	Host        string   `json:"host,omitempty"`
	Port        string   `json:"port,omitempty"`
	DBName      string   `json:"dbName,omitempty"`
	UserName    string   `json:"userName,omitempty"`
	Password    string   `json:"password,omitempty"`
	MinPoolSize int      `json:"minPoolSize,omitempty"`
	MaxPoolSize int      `json:"maxPoolSize,omitempty"`
	Connection  *gorm.DB `json:"connection,omitempty"`
}

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
	x, err := gorm.Open("postgres", pgInfo)
	if err != nil {
		log.Fatalf("Fail to connect pg: %v\n", err)
	}

	x.LogMode(true)
	x.SingularTable(true)
	x.DB().SetMaxIdleConns(configs.Cf.Db.Postgres.MinPoolSize)
	x.DB().SetMaxOpenConns(configs.Cf.Db.Postgres.MaxPoolSize)

	initTable(x)
	initData()

	PGDB = PGService{
		Host:        configs.Cf.Db.Postgres.Host,
		Port:        configs.Cf.Db.Postgres.Port,
		DBName:      configs.Cf.Db.Postgres.DBName,
		UserName:    configs.Cf.Db.Postgres.Username,
		Password:    configs.Cf.Db.Postgres.Password,
		MinPoolSize: configs.Cf.Db.Postgres.MinPoolSize,
		MaxPoolSize: configs.Cf.Db.Postgres.MaxPoolSize,
		Connection:  x,
	}
}

func initTable(x *gorm.DB) {
	x.AutoMigrate(&models.User{})
	x.AutoMigrate(&models.DailyNote{})
}

func initData() {
}
