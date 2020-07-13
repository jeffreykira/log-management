package pkg

import (
	"os"
	"time"

	log "github.com/mgutz/logxi/v1"
)

var timeLogger log.Logger

func init() {
	timeLogger = log.NewLogger(log.NewConcurrentWriter(os.Stdout), "time")
}

//CurrentTime ...
func CurrentTime() string {
	timeLogger.Debug("CurrentTime start")

	t := time.Now()
	t.Format("1988-09-18 23:00:00")
	timeLogger.Info("CurrentTime", "get time", t.String())

	return t.String()
}
