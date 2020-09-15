package configs

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var (
	Cf *Conf
)

func InitConfig() {
	log.Info("starting init config")
	if isGetConfigFromEnv() {
		Cf = getConfigFromEnv()
	} else {
		Cf = getConf()
	}

	getExtraOsEnv()

	log.Infof("deploy: %s", Cf.Deploy)
}

type db struct {
	Postgres postgres `yaml:"postgres" required:"true"`
}

type postgres struct {
	Host        string `yaml:"host" required:"true"`
	Username    string `yaml:"username" required:"true"`
	Port        string `yaml:"port" required:"true"`
	Password    string `yaml:"password" required:"true"`
	DBName      string `yaml:"dbname" required:"true"`
	MinPoolSize int    `yaml:"minPoolSize" required:"true" default:"10"`
	MaxPoolSize int    `yaml:"maxPoolSize" required:"true" default:"30"`
}

type server struct {
	Port   string `yaml:"port" required:"true"`
	Domain string `yaml:"domain" required:"true"`
}

type jwt struct {
	ExpiredDays int    `yaml:"expiredDays" required:"true"`
	Secret      string `yaml:"secret" required:"true"`
}

type Conf struct {
	Db     db     `yaml:"db" required:"true"`
	Server server `yaml:"server" required:"true"`
	Deploy string `yaml:"deploy" required:"true"`
	Jwt    jwt    `yaml:"jwt" required:"true"`
}

func getConf() *Conf {
	c := &Conf{}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf(err.Error())
	}
	yamlFile, err := ioutil.ReadFile(dir + "/config-dev.yml")
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return c
}

func isGetConfigFromEnv() bool {
	deploy := os.Getenv("DEPLOY")
	if deploy == "k8s" {
		return true
	}
	return false
}

func getConfigFromEnv() *Conf {
	cf := &Conf{}
	err := envconfig.Process("", cf)
	if err != nil {
		log.Fatalf("get env error, " + err.Error())
	}
	getExtraParamsFromEnv(cf)

	return cf
}

func getExtraParamsFromEnv(cf *Conf) {
}

func getExtraOsEnv() {
}
