package gateways

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h HTTPGateway) GetAllUsers(c echo.Context) error {
	res, err := h.APIService.GetAllUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string][]string{
			"error": {err.Error()},
		})
	}

	return c.JSON(http.StatusOK, res)
}
