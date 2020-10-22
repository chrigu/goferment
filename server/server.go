package server

import (
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func router(ch chan string) http.Handler {
	router := gin.Default()
	router.Use(gin.Recovery())

	if os.Getenv("PROXY_TO_VUE_DEV_SERVER") == "true" {
		router.NoRoute(reverseProxy())
	}

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			ch <- "on"
			c.JSON(
				http.StatusOK,
				gin.H{
					"code":  http.StatusOK,
					"error": "Welcome server 01",
				},
			)
		})

		api.GET("/test", func(c *gin.Context) {
			ch <- "off"
			c.JSON(
				http.StatusOK,
				gin.H{
					"code":  http.StatusOK,
					"error": "Welcome server 01",
				},
			)
		})
	}

	return router
}

// Inits webserver
func CreateServer(ch chan string) *http.Server {
	server := &http.Server{
		Addr:         ":8000",
		Handler:      router(ch),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server
}

// https://github.com/gin-gonic/gin/issues/686
func reverseProxy() gin.HandlerFunc {

	target := "localhost:8080"

	return func(c *gin.Context) {
		director := func(req *http.Request) {
			// r := c.Request
			//req = r
			req.URL.Scheme = "http"
			req.URL.Host = target
			// req.Header["my-header"] = []string{r.Header.Get("my-header")}
			// Golang camelcases headers
			delete(req.Header, "My-Header")
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
