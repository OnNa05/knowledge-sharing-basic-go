package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/OnNa05/knowledge-sharing-basic-go/mongodb/repositories"
	"github.com/OnNa05/knowledge-sharing-basic-go/scr/auth/dao"
)

type AuthService struct {
	UserRepo repositories.IUserRepo
}

type IAuthService interface {
	Register(ctx context.Context, d dao.RegisterRequest) error
	Login(ctx context.Context, d dao.LoginRequest) (dao.LoginResponse, error)
}

func NewAuthenService(a repositories.IUserRepo) IAuthService {
	return &AuthService{
		UserRepo: a,
	}
}

func (sv AuthService) NewSHA256(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}
