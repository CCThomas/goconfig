package web

import (
  "io/ioutil"
  "net/http"

  "github.com/gin-gonic/gin"
  log "github.com/sirupsen/logrus"
)

func GetJson(c *gin.Context) ([]byte, error) {
  log.Info("Getting Json from Request")
  requestJson, err := ioutil.ReadAll(c.Request.Body)
  if err != nil {
    log.Error(err)
    c.IndentedJSON(http.StatusBadRequest, "Request body json format is invalid.")
    return nil, err
  }
  return requestJson, nil
}
