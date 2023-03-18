package services

import (
	"context"

	"github.com/OnNa05/knowledge-sharing-basic-go/scr/user/dao"
)

func (s *APIService) GetAllUsers(
	ctx context.Context,
) ([]dao.User, error) {
	result := make([]dao.User, 0)
	entities, err := s.UserRepo.FindAll(ctx)
	if err != nil {
		return result, err
	}

	for _, entity := range entities {
		result = append(result, dao.User{
			ID:       entity.ID,
			Name:     entity.Name,
			Email:    entity.Email,
			Password: entity.Password,
			Role:     entity.Role,
			CreateAt: entity.CreateAt,
		})
	}

	return result, nil
}
