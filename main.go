package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

// add logging middleware
// log information about the request
func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		startTime := time.Now()

		err := next(c)
		if err != nil {
			c.Error(err)
		}

		// Log info
		endTime := time.Now()
		_ = endTime.Sub(startTime) // duration
		clientIP := c.RealIP()
		method := c.Request().Method
		path := c.Path()
		statusCode := c.Response().Status

		fmt.Printf("[%s] %s %s %d %s\n", endTime.Format(time.RFC3339), clientIP, method, statusCode, path)
		// fmt.Printf("Request processing time: %s\n", duration)

		return nil
	}
}

// add authentication middleware
func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing Authorization header")
		}

		authParts := strings.Split(authHeader, " ")

		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid Authorization header format")
		}

		token := authParts[1]

		// token validation logic
		if token != "token" {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		return next(c)
	}
}

func main() {
	e := echo.New()

	e.Use(LoggingMiddleware)

	e.Use(AuthenticationMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/protected", func(c echo.Context) error {
		return c.String(http.StatusOK, "Protected resource")
	})

	e.Logger.Fatal((e.Start(":1323")))
}
