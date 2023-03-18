package gateways

import (
	"net/http"

	"github.com/OnNa05/knowledge-sharing-basic-go/scr/user/dao"

	"github.com/labstack/echo"
)

func (h HTTPGateway) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()
	user := new(dao.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	err := h.APIService.DeleteUser(ctx, dao.DeleteUserRequest{
		ID: user.ID,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"msg": "Delete Success!!",
	})
}
