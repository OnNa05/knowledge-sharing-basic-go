package gateways

import (
	"net/http"

	"github.com/OnNa05/knowledge-sharing-basic-go/scr/auth/dao"
	"github.com/labstack/echo/v4"
)

func (h HTTPGateway) Login(c echo.Context) error {
	var r dao.LoginRequest
	err := c.Bind(&r)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	res, err := h.AuthenService.Login(c.Request().Context(), r)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}
