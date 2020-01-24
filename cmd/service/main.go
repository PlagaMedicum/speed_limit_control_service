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
    "Access-Control-Allow-Origin":  "*",
    "Access-Control-Allow-Headers": "access-control-allow-origin, content-type",
    "Content-Type":                 "application/json",
}

func httpMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        for k, v := range httpHeaders {
            c.Set(k, v)
        }

        if c.Request.Method == http.MethodOptions {
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
    r.Use(gin.Logger())
    r.Use(httpMiddleware())

    r.POST("/add", h.AddDataHandler)
    r.GET("/infractions", h.GetInfractionsHandler)
    r.GET("/boundaries", h.GetMinMaxHandler)
    r.OPTIONS("/add", func(c *gin.Context){})

    err := r.Run(":8080")
    if err != nil {
        log.Fatal("Server crashed: ", err)
    }
}
