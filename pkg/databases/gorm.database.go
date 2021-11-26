package databases

import (
  "fmt"
  "os"

  log "github.com/sirupsen/logrus"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var GormDB *gorm.DB

func ConfigureGorm() {
  log.Info("Try to connect to database.")
  host := os.Getenv("GOCONFIG_DB_HOST")
  user := os.Getenv("GOCONFIG_DB_USER")
  password := os.Getenv("GOCONFIG_DB_PASSWORD")
  dbname := os.Getenv("GOCONFIG_DB_NAME")
  port := os.Getenv("GOCONFIG_DB_PORT")
  sslmode := os.Getenv("GOCONFIG_DB_SSLMODE")

  dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Chicago",
    host, user, password, dbname, port, sslmode)
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatal("failed to connect database")
  }

  log.Infof("Connected to %s.\n", host)
  GormDB = db
}
