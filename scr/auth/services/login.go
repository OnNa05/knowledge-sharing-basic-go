package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/OnNa05/knowledge-sharing-basic-go/scr/auth/dao"
	"github.com/dgrijalva/jwt-go"
)

func (sv AuthService) Login(ctx context.Context, d dao.LoginRequest) (dao.LoginResponse, error) {
	email := strings.ToLower(d.Email)
	password := sv.NewSHA256(d.Password)

	u, err := sv.UserRepo.FindByEmailAndPassword(ctx, email, password)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return dao.LoginResponse{}, fmt.Errorf("email or password incorrect please try again")
		}
		return dao.LoginResponse{}, err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	cl := token.Claims.(jwt.MapClaims)
	now := time.Now().Unix()
	cl["access"] = now
	cl["identity"] = u.ID.Hex()
	cl["type"] = "access"

	t, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		return dao.LoginResponse{}, err
	}

	_, err = sv.UserRepo.UpdateToken(ctx, email, t)
	if err != nil {
		return dao.LoginResponse{}, err
	}

	u, err = sv.UserRepo.FindByEmail(ctx, u.Email)

	return dao.LoginResponse{
		Message: "success",
		Token:   t,
	}, err
}
