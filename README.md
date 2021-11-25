# goconfig
Golang Configuration Package to assist with developing simple micro services.

## Version
github.com/CCThomas/goconfig v0.1

## Requirements
- Golang verison [1.6](https://golang.org/doc/go1.6)

## Usage

### Envrionemnt Variables
| Variable | Notes |
| --- | --- |
| `GOCONFIG_LOG_FILE_NAME` | Name of log file. |
| `GOCONFIG_LOG_LEVEL` | Log levels can be found [here](https://github.com/sirupsen/logrus#level-logging). |
| `GOCONFIG_LOG_TO_TERMINAL` | If `true `, logs will be printed to the terminal  |

### Logger
```go
import "github.com/CCThomas/goconfig/pkg/loggers"
loggers.ConfigureLogrus()
```
```env
GOCONFIG_LOG_FILE_NAME=my-service.log
GOCONFIG_LOG_LEVEL=debug
GOCONFIG_LOG_TO_TERMINAL=true
```

## Third Party
- [github.com/joho/godotenv v1.4.0](https://github.com/joho/godotenv) - MIT License
    - Go (golang) port of the Ruby dotenv project (which loads env vars from a .env file)
- [github.com/sirupsen/logrus v1.8.1](https://github.com/sirupsen/logrus) - MIT License
    - Logrus is a structured logger for Go (golang), completely API compatible with the standard library logger.
