package gateways

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func (h HTTPGateway) GetUser(c echo.Context) error {
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(jwt.MapClaims)
	uid := claims["identity"].(string)

	res, err := h.APIService.GetUser(c.Request().Context(), uid)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string][]string{
			"error": {err.Error()},
		})
	}

	return c.JSON(http.StatusOK, res)
}
