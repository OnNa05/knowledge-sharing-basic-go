package gateways

import (
	services "github.com/OnNa05/knowledge-sharing-basic-go/scr/auth/service"
	"github.com/labstack/echo"
)

type HTTPGateway struct {
	AuthenService services.IAuthService
}

func NewHTTPGateway(g *echo.Group, sv services.IAuthService) {
	h := &HTTPGateway{
		AuthenService: sv,
	}

	g.POST("/register", h.Register)
	g.POST("/login", h.Login)
}
