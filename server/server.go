package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func router(ch chan string) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		ch <- "on"
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})

	e.GET("/test", func(c *gin.Context) {
		ch <- "off"
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})

	return e
}

// Inits webserver
func CreateServer(ch chan string) *http.Server {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router(ch),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server
}
