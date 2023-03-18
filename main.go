package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// add logging middleware
// log information about the request
func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		startTime := time.Now()

		// Call the next handler in the chain
		err := next(c)
		if err != nil {
			c.Error(err)
		}

		// Log info
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		clientIP := c.RealIP()
		method := c.Request().Method
		path := c.Path()
		statusCode := c.Response().Status

		fmt.Printf("[%s] %s %s %d %s\n", endTime.Format(time.RFC3339), clientIP, method, statusCode, path)
		fmt.Printf("Request processing time: %s\n", duration)
		
		return nil
	}
}

func main() {
	e := echo.New()

	e.Use(LoggingMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal((e.Start(":1323")))
}
