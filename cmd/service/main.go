package main

import (
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/handlers"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/repositories"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/usecases"
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
    "net/http"
)

var httpHeaders = map[string]string{
    "Access-Control-Allow-Headers": "Content-Type, api_key, Authorization, access-control-allow-origin",
    "Access-Control-Allow-Origin":  "*",
    "Content-Type":                 "application/json",
    "Access-Control-Allow-Methods": "POST, OPTIONS, GET",
}

func httpMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        for k, v := range httpHeaders {
            c.Writer.Header().Set(k, v)
        }

        if c.Request.Method == http.MethodOptions {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    }
}



func main() {
    h := handlers.Controller{
        usecases.Controller{
            repositories.Controller{
                // TODO
            },
        },
    }

    r := gin.New()
    r.Use(gin.Logger()).Use(httpMiddleware())

    r.POST("/add", h.AddDataHandler)
    r.GET("/infractions", h.GetInfractionsHandler)
    r.GET("/boundaries", h.GetMinMaxHandler)

    err := r.Run(":8080")
    if err != nil {
        log.Fatal("Server crashed: ", err)
    }
}
