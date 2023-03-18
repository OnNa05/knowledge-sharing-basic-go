package gateways

import (
	"github.com/OnNa05/knowledge-sharing-basic-go/scr/user/services"
	"github.com/labstack/echo"
)

type HTTPGateway struct {
	APIService services.IAPIService
}

func NewHTTPGateway(g *echo.Group, sv services.IAPIService) {
	h := &HTTPGateway{
		APIService: sv,
	}

	g.GET("/", h.GetAllUsers)
	g.POST("/", h.CreateUser)
	g.PUT("/", h.UpdateUser)
	g.DELETE("/", h.DeleteUser)
}
