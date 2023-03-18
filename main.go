package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/OnNa05/knowledge-sharing-basic-go/mongodb/connection"
	repo "github.com/OnNa05/knowledge-sharing-basic-go/mongodb/repositories"
	gw "github.com/OnNa05/knowledge-sharing-basic-go/scr/user/gateways"
	sv "github.com/OnNa05/knowledge-sharing-basic-go/scr/user/services"

	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
)

// add custom logging middleware
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

	e.Use(middleware.Logger())

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!")
	})

	m := connection.NewMongoDB(os.Getenv("MONGODB_URI"))

	repo0 := repo.NewUserRepo(m)

	apiSV := sv.NewAPIService(repo0)

	gw.NewHTTPGateway(e.Group(""), apiSV)

	e.Logger.Fatal((e.Start(":1323")))
}
