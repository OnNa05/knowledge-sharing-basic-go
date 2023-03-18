package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/OnNa05/knowledge-sharing-basic-go/mongodb/connection"
	repo "github.com/OnNa05/knowledge-sharing-basic-go/mongodb/repositories"
	gw "github.com/OnNa05/knowledge-sharing-basic-go/scr/user/gateways"
	sv "github.com/OnNa05/knowledge-sharing-basic-go/scr/user/services"
	"github.com/dgrijalva/jwt-go"

	authgw "github.com/OnNa05/knowledge-sharing-basic-go/scr/auth/gateways"
	authsv "github.com/OnNa05/knowledge-sharing-basic-go/scr/auth/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
func UIDHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		rawU := c.Get("user")
		if rawU == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "session is expired.",
			})
		}

		u := rawU.(*jwt.Token)
		claims := u.Claims.(jwt.MapClaims)
		if claims["identity"] == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "session is expired.",
			})
		}

		uid := claims["identity"].(string)
		fmt.Println("UIDHeader UID: ", uid)
		c.Request().Header.Set("bn-uid", uid)
		return next(c)
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!")
	})

	m := connection.NewMongoDB("")

	repo0 := repo.NewUserRepo(m)

	apiSV := sv.NewAPIService(repo0)
	authSrv := authsv.NewAuthenService(repo0)

	userrg := e.Group("/user")
	jwtRequiredRoute := userrg.Group(
		"",
		middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey: []byte("SECRET"),
		}),
		UIDHeader,
	)
	userg := jwtRequiredRoute.Group(
		"",
	)

	gw.NewHTTPGateway(userg, apiSV)
	authgw.NewHTTPGateway(e.Group("/auth"), authSrv)

	e.Logger.Fatal((e.Start(":1323")))
}
