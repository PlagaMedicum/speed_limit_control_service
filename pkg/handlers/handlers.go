package handlers

import (
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/domain"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/usecases"
    "github.com/gin-gonic/gin"
    "github.com/pkg/errors"
    log "github.com/sirupsen/logrus"
    "net/http"
    "strconv"
    "time"
)

const dateLayout string = "02.01.2006 15:04:05"

type Controller struct {
    usecases.Usecases
}

type jsonMessage struct {
    Date string
    Number string
    Speed  float32
}

func unmarshalJSONMessage(m jsonMessage) (domain.SpeedInfo, error) {
    si := domain.SpeedInfo{
        Number: m.Number,
        Speed: m.Speed,
    }

    var err error
    si.Date, err = time.Parse(dateLayout, m.Date)
    return si, errors.Wrap(err, "Error parsing date")
}

func marshalJSONMessage(si domain.SpeedInfo) jsonMessage {
    return jsonMessage{
        Date:   si.Date.Format(dateLayout),
        Number: si.Number,
        Speed:  si.Speed,
    }
}

// AddDataHandler is http handler that adds new speed information in the storage.
func (ctl Controller) AddDataHandler(c *gin.Context) {
    var m jsonMessage
    err := c.ShouldBindJSON(&m)
    if err != nil {
        err = errors.Wrap(err, "Error decoding request body")
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        log.Error(err)
    }

    si, err := unmarshalJSONMessage(m)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        log.Error(err)
    }

    err = ctl.Usecases.AddData(c, si)
    if err != nil {
        err = errors.Wrap(err, "Error adding data")
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        log.Error(err)
    }
    c.Status(http.StatusCreated)
}

// GetInfractionsHandler is http handler that returns a list of all transport
// that broke provided speed limit, for the specified date.
func (ctl Controller) GetInfractionsHandler(c *gin.Context) {
    date, err := time.Parse(dateLayout, c.Query("date"))
    if err != nil {
        err = errors.Wrap(err, "Error parsing date from query")
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        log.Error(err)
    }

    limit, err := strconv.ParseFloat(c.Query("speed"), 32)
    if err != nil {
        err = errors.Wrap(err, "Error parsing speed(float32) from query")
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        log.Error(err)
    }

    silist, err := ctl.GetInfractions(c, date, float32(limit))
    if err != nil {
        err = errors.Wrap(err, "Error getting data")
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        log.Error(err)
    }

    var resp []jsonMessage
    for _, si := range silist {
        resp = append(resp, marshalJSONMessage(si))
    }

    c.JSON(http.StatusOK, resp)
}

// GetMinMaxHandler is http handler that returns minimal and maximal recorded speeds
// for specified date.
func (ctl Controller) GetMinMaxHandler(c *gin.Context) {
    date, err := time.Parse(dateLayout, c.Query("date"))
    if err != nil {
        err = errors.Wrap(err, "Error parsing date from query")
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        log.Error(err)
    }

    silist, err := ctl.GetMinMax(c, date)
    if err != nil {
        err = errors.Wrap(err, "Error getting data")
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        log.Error(err)
    }

    var resp []jsonMessage
    for _, si := range silist {
        resp = append(resp, marshalJSONMessage(si))
    }

    c.JSON(http.StatusOK, resp)
}
