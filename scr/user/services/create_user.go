package services

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/OnNa05/knowledge-sharing-basic-go/mongodb/entities"
	"github.com/OnNa05/knowledge-sharing-basic-go/scr/user/dao"
)

func (s *APIService) CreateUser(
	ctx context.Context,
	d dao.User,
) (dao.CreateUserResponse, error) {

	ent := &entities.UserEntitty{
		Name:     d.Name,
		Email:    d.Email,
		Password: d.Password,
		Role:     d.Role,
		CreateAt: d.CreateAt,
	}
	res, err := s.UserRepo.Insert(ctx, ent)
	if err != nil {
		return dao.CreateUserResponse{}, err
	}

	return dao.CreateUserResponse{
		ID: res.InsertedID.(primitive.ObjectID).Hex(),
	}, err
}
