# log-management
Use [logxi](https://github.com/mgutz/logxi) library


## Getting Started
logxi can set unique identifier for log，並透過環境變數來控制 log level 以及要顯示的 log identifier
```
import log "github.com/mgutz/logxi/v1"

var logger log.Logger

func init() {
    // create a logger with a unique identifier
	logger = log.NewLogger(log.NewConcurrentWriter(os.Stdout), "main")
}

func main() {
    logger.Debug("main start")

    logger.Info("process done.")
}
```


## Level Map
- Fatal: FTL
- Error: ERR
- Warn : WRN
- Info : INF
- Debug: DBG
- Trace: TRC


## Log Configuration
#### defaults to showing warnings and above
```
go run cmd/main.go
```
#### show all logs
```
LOGXI=* go run cmd/main.go
```
#### time show debug, other show error log
```
LOGXI=*=ERR,time=DBG go run cmd/main.go
```
#### time close log, other show debug log
```
LOGXI=*=DBG,time=OFF go run cmd/main.go
```
#### set all to Error and set dat related packages to Debug
```
LOGXI=*=ERR,dat*=DBG go run cmd/main.go
```
