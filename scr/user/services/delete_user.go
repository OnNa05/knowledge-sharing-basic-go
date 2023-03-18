package services

import (
	"context"

	"github.com/OnNa05/knowledge-sharing-basic-go/scr/user/dao"
)

func (s *APIService) DeleteUser(ctx context.Context, d dao.DeleteUserRequest) error {
	_, err := s.UserRepo.DeleteUserByID(ctx, d.ID)
	if err != nil {
		return err
	}

	return err
}
