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

### Web
Get Json
```go
import (
  "github.com/CCThomas/goconfig/pkg/web"
  "github.com/gin-gonic/gin"
)
// c *gin.Context
requestJson, err := web.GetJson(c)
if err != nil {
  // Response has already been set.
  return
}
```

## Third Party
- [github.com/gin-gonic/gin v1.7.7](https://github.com/gin-gonic/gin) - MIT License
    - Gin is a web framework written in Go (Golang).
- [github.com/joho/godotenv v1.4.0](https://github.com/joho/godotenv) - MIT License
    - Go (golang) port of the Ruby dotenv project (which loads env vars from a .env file)
- [github.com/sirupsen/logrus v1.8.1](https://github.com/sirupsen/logrus) - MIT License
    - Logrus is a structured logger for Go (golang), completely API compatible with the standard library logger.
- [gopkg.in/go-playground/assert.v1 v1.2.1](https://github.com/go-playground/assert) - MIT License
    - Package assert is a Basic Assertion library used along side native go testing
- [gopkg.in/go-playground/validator.v9 v9.29.1](https://github.com/go-playground/validator/) - MIT License
    - Package validator implements value validations for structs and individual fields based on tags.
- [gorm.io/driver/postgres v1.2.2](https://github.com/go-gorm/postgres) - MIT License
    - postgres's database/sql driver for Gorm
- [gorm.io/gorm v1.22.3](https://gorm.io) - MIT License
    - The fantastic ORM library for Golang, aims to be developer friendly.
