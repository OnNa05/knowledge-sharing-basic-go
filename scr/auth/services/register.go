package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/OnNa05/knowledge-sharing-basic-go/mongodb/entities"
	"github.com/OnNa05/knowledge-sharing-basic-go/scr/auth/dao"
)

func (sv AuthService) Register(ctx context.Context, d dao.RegisterRequest) error {
	email := strings.ToLower(d.Email)
	password := sv.NewSHA256(d.Password)

	res, _ := sv.UserRepo.FindByEmail(ctx, email)
	if res != nil {
		errors.New("email already")
	}

	ent := entities.UserEntitty{
		Name:     d.Name,
		Email:    email,
		Password: password,
		Role:     "user",
		CreateAt: time.Now(),
	}

	_, err := sv.UserRepo.Insert(ctx, &ent)
	if err != nil {
		return err
	}

	return nil
}
