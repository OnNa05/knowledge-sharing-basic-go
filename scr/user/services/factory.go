package services

import (
	"context"

	"github.com/OnNa05/knowledge-sharing-basic-go/mongodb/entities"
	"github.com/OnNa05/knowledge-sharing-basic-go/mongodb/repositories"

	"github.com/OnNa05/knowledge-sharing-basic-go/scr/user/dao"
)

type (
	APIService struct {
		UserRepo repositories.IUserRepo
	}

	IAPIService interface {
		GetAllUsers(ctx context.Context) ([]dao.User, error)
		CreateUser(ctx context.Context, User dao.User) (dao.CreateUserResponse, error)
		UpdateUser(ctx context.Context, User dao.UpdateUserRequest) (dao.UpdateUserResponse, error)
		DeleteUser(ctx context.Context, User dao.DeleteUserRequest) error
		GetUser(ctx context.Context, id string) (entities.UserEntitty, error)
	}
)

func NewAPIService(
	repo0 repositories.IUserRepo,
) IAPIService {
	return &APIService{
		UserRepo: repo0,
	}
}
