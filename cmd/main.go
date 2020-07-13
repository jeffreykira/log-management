package main

import (
	"log-example/pkg"
	"os"

	log "github.com/mgutz/logxi/v1"
)

var logger log.Logger

func init() {
	logger = log.NewLogger(log.NewConcurrentWriter(os.Stdout), "main")
}

func main() {
	t := pkg.CurrentTime()
	logger.Debug("get current time", "time", t)

	err := pkg.DatabaseSave(t)
	if err != nil {
		logger.Error("panic", "message", err)
	}

	logger.Info("process done.")
}
