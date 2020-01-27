package main

import (
    "fmt"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/storage"
    "github.com/pkg/errors"
    "net/http"
    "time"

    "github.com/PlagaMedicum/speed_limit_control_service/pkg/handlers"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/repositories"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/usecases"
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
)

type config struct {
    startTime time.Time
    endTime time.Time
    httpHeaders map[string]string
    dataPath string
}

const accessTimeLayout = "15:04:05"

func clockIsBetween(start, end, check time.Time) bool {
    return check.After(start) && check.Before(end)
}

func httpMiddleware(cfg config) gin.HandlerFunc {
	return func(c *gin.Context) {
		for k, v := range cfg.httpHeaders {
			c.Writer.Header().Set(k, v)
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		if c.Request.Method == http.MethodGet {
            now := time.Now()
            check, err := time.Parse(accessTimeLayout, fmt.Sprintf("%02d:%02d:%02d", now.Hour(), now.Minute(), now.Second()))
            if err != nil {
                log.Errorf("Error parsing current time on the server: %v", err)
            }

            if !clockIsBetween(cfg.startTime, cfg.endTime, check) {
                err := errors.Errorf("Error accessing service storage. Server provides storage access from" +
                    "%02d:%02d to %02d:%02d, but the time now is %02d:%02d:%02d",
                    cfg.startTime.Hour(), cfg.startTime.Minute(), cfg.endTime.Hour(), cfg.endTime.Minute(),
                    now.Hour(), now.Minute(), now.Second())
                c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            }
        }

		c.Next()
	}
}

func parseConfigurations() (config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()
    if err != nil {
        return config{}, errors.Wrap(err, "Error reading config file")
    }
    cfg := config{}

    cfg.startTime, err = time.Parse(accessTimeLayout, viper.GetString("startTime"))
    if err != nil {
        return config{}, err
    }
    cfg.endTime, err = time.Parse(accessTimeLayout, viper.GetString("endTime"))
    if err != nil {
        return config{}, err
    }

    cfg.httpHeaders = viper.GetStringMapString("cors")

    cfg.dataPath = viper.GetString("dataPath")

    return cfg, nil
}

func main() {
    cfg, err := parseConfigurations()
    if err != nil {
        log.Fatalf("Error parsing configurations: %v", err)
    }

	h := handlers.Controller{
		usecases.Controller{
			repositories.Controller{
			    storage.Storage{
			        DataPath: cfg.dataPath,
                },
            },
		},
	}

	r := gin.New()
	r.Use(gin.Logger()).Use(httpMiddleware(cfg))

	r.POST("/add", h.AddDataHandler)
	r.GET("/infractions", h.GetInfractionsHandler)
	r.GET("/boundaries", h.GetMinMaxHandler)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Server crashed: ", err)
	}
}
