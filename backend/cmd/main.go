package main

import (
	"os"
	"time"

	"backend/logs"
	"backend/utils"
	"backend/webserver"
	"go.uber.org/zap"
)

var (
	logger *zap.SugaredLogger
)

type config struct {
	WebServer *webserver.Config `json:"webserver"`
}

func main() {
	logger = logs.GetLogger()

	cfgPath := "config.json"
	if len(os.Args) > 1 {
		cfgPath = os.Args[1]
	}

	var cfg config
	err := utils.ReadConfig(cfgPath, &cfg)
	if err != nil {
		logger.Fatalf("failed to read config: %s", err)
		return
	}

	forHTY(cfg)
}

func forHTY(cfg config) {
	if cfg.WebServer != nil {
		logger.Info("manager server starting ...")
		go func() {
			err := webserver.Run(*cfg.WebServer)
			if err != nil {
				logger.Fatalf("webserver run exit: %s", err)
				return
			}
		}()
	}

	for {
		<-time.After(time.Hour)
		logger.Debug("An Hour Passed...")
	}
}
