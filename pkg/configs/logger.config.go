package configs

import (
  "io"
  "os"

  "github.com/joho/godotenv"
  log "github.com/sirupsen/logrus"
)

func ConfigureLogrus() {
  // Log as JSON instead of the default ASCII formatter.
  log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

  // Output to stdout instead of the default stderr
  // Can be any io.Writer, see below for File example
  log.SetOutput(os.Stdout)

  // Only log the warning severity or above.
  logLevelStr := os.Getenv("LOG_LEVEL")
  logLevel := log.ErrorLevel
  if logLevelStr != "" {
    var err error
    logLevel, err = log.ParseLevel(logLevelStr)
    if err != nil {
      panic("invalid environment variable value set for: LOG_LEVEL")
    }
  }

  log.SetLevel(logLevel)

  // Configure Logger.
  logFileName := os.Getenv("LOG_FILE_NAME")
  if logFileName == "" {
    logFileName = "logfile.log"
  }

  file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
  if err != nil {
    log.Fatal(err)
  }

  // Configure if we should log to terminal.
  if os.Getenv("LOG_TO_TERMINAL") == "true" {
    multi := io.MultiWriter(file, os.Stdout)
    log.SetOutput(multi)
  }

  // Load the .env file in the current directory
  log.Infof("Loading .env file.")
  godotenv.Load()
}
