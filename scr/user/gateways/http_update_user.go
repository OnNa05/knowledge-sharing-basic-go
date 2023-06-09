package gateways

import (
	"net/http"

	"github.com/OnNa05/knowledge-sharing-basic-go/scr/user/dao"
	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
)

func (h HTTPGateway) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()
	user := new(dao.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	_, err = h.APIService.UpdateUser(ctx, dao.UpdateUserRequest{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword),
		Role:     user.Role,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"msg": "Update Success!!",
	})
}
