package gateways

import (
	"fmt"
	"net/http"

	"github.com/OnNa05/knowledge-sharing-basic-go/scr/auth/dao"
	"github.com/labstack/echo"
)

func (h HTTPGateway) Register(c echo.Context) error {
	var r dao.RegisterRequest
	err := c.Bind(&r)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	err = h.AuthenService.Register(c.Request().Context(), r)
	if err != nil {
		if err.Error() == "email already" {
			return c.JSON(http.StatusConflict, map[string]string{
				"error": fmt.Sprintf("email %v is already exist", r.Email),
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "register successful",
	})
}
