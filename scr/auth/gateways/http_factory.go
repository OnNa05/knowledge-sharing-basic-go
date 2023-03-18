package gateways

import (
	service "github.com/OnNa05/knowledge-sharing-basic-go/scr/auth/services"
	"github.com/labstack/echo"
)

type HTTPGateway struct {
	AuthenService service.IAuthService
}

func NewHTTPGateway(g *echo.Group, sv service.IAuthService) {
	h := &HTTPGateway{
		AuthenService: sv,
	}

	g.POST("/register", h.Register)

}
