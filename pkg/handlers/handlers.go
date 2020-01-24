package handlers

import (
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/domain"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/usecases"
    "github.com/gin-gonic/gin"
    "github.com/pkg/errors"
    log "github.com/sirupsen/logrus"
    "net/http"
)

type Controller struct {
    usecases.Usecases
}

// AddDataHandler is http handler that adds new speed information in the storage.
func (con Controller) AddDataHandler(c *gin.Context) {
    var si domain.SpeedInfo
    err := c.ShouldBindJSON(&si)
    if err != nil {
        err = errors.Wrap(err, "Error decoding request body")
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        log.Error(err)
    }

    err = con.Usecases.AddData(c, si)
    if err != nil {
        err = errors.Wrap(err, "Error adding data")
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        log.Error(err)
    }
    c.Status(http.StatusCreated)
}

// GetInfractionsHandler is http handler that returns a list of all transport
// that broke provided speed limit, for the specified date.
func (con Controller) GetInfractionsHandler(c *gin.Context) {

    c.Status(http.StatusOK)
}

// GetMinMaxHandler is http handler that returns minimal and maximal recorded speeds
// for specified date.
func (con Controller) GetMinMaxHandler(c *gin.Context) {

    c.Status(http.StatusOK)
}
