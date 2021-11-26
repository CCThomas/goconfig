# goconfig
Golang Configuration Package to assist with developing simple micro services.

## Requirements
- Golang verison [1.6](https://golang.org/doc/go1.6)

## Usage

### Envrionemnt Variables
| Variable | Notes |
| --- | --- |
| `GOCONFIG_DB_HOST` | Database Host. Example: `postgres` |
| `GOCONFIG_DB_NAME` | Database Name |
| `GOCONFIG_DB_PASSWORD` | Database Password |
| `GOCONFIG_DB_PORT` | Database Port |
| `GOCONFIG_DB_SSLMODE` | Should Database use SSLMode |
| `GOCONFIG_DB_USER` | Database User |
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

### Database
**Currently Postgres is the only supported database**
```go
import "github.com/CCThomas/goconfig/pkg/databases"
databases.ConfigureGorm()
database.DB.Create(&myModel)
```
```env
GOCONFIG_DB_HOST=localhost
GOCONFIG_DB_USER=myuser
GOCONFIG_DB_PASSWORD=mypassword
GOCONFIG_DB_NAME=dbname
GOCONFIG_DB_PORT=5434
GOCONFIG_DB_SSLMODE=disable
```

## Third Party
- [github.com/joho/godotenv v1.4.0](https://github.com/joho/godotenv) - MIT License
    - Go (golang) port of the Ruby dotenv project (which loads env vars from a .env file)
- [github.com/sirupsen/logrus v1.8.1](https://github.com/sirupsen/logrus) - MIT License
    - Logrus is a structured logger for Go (golang), completely API compatible with the standard library logger.
- [gorm.io/driver/postgres v1.2.2](https://github.com/go-gorm/postgres) - MIT License
    - postgres's database/sql driver for Gorm
- [gorm.io/gorm v1.22.3](https://gorm.io) - MIT License
    - The fantastic ORM library for Golang, aims to be developer friendly.
