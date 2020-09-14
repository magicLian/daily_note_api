package crons

import (
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

const (
	cleanPVDuration                   = "0 0 2 * * ?"
)

func InitCrons() {
	log.Info("starting crons...")
	crons := cron.New()

	crons.Start()
}
