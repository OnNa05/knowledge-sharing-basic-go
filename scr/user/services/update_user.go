package services

import (
	"context"

	"github.com/OnNa05/knowledge-sharing-basic-go/scr/user/dao"
)

func (s *APIService) UpdateUser(ctx context.Context, d dao.UpdateUserRequest) (dao.UpdateUserResponse, error) {
	user, err := s.UserRepo.FindByID(ctx, d.ID.Hex())
	if err != nil {
		return dao.UpdateUserResponse{}, err
	}

	user.Name = d.Name
	user.Email = d.Email
	user.Password = d.Password
	user.Role = d.Role

	_, err = s.UserRepo.UpdateUser(ctx, user)
	if err != nil {
		return dao.UpdateUserResponse{}, err
	}

	return dao.UpdateUserResponse{
		ID: user.ID.Hex(),
	}, err
}
