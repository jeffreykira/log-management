package pkg

import (
	"database/sql"
	"os"

	log "github.com/mgutz/logxi/v1"
)

var dbLogger log.Logger

func init() {
	dbLogger = log.NewLogger(log.NewConcurrentWriter(os.Stdout), "database")
}

//DatabaseSave ...
func DatabaseSave(currentTime string) error {
	dbLogger.Debug("Save", "input", currentTime)

	_, err := sql.Open("postgres", "dbname=testdb")
	if err != nil {
		dbLogger.Warn("Could not open database", "err", err)
		return err
	}

	return nil
}
