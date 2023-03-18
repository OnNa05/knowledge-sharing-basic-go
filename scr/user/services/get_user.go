package services

import (
	"context"

	"github.com/OnNa05/knowledge-sharing-basic-go/mongodb/entities"
)

func (s *APIService) GetUser(ctx context.Context, id string) (entities.UserEntitty, error) {
	res, err := s.UserRepo.FindByID(ctx, id)
	if err != nil {
		return entities.UserEntitty{}, err
	}

	return *res, nil
}
