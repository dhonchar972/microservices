package author

import (
	"context"
	"restApi/pkg/logging"
)

/*
User repository.
*/
type Service struct {
	repository Repository
	logger     *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateAuthorDto) (Author, error) {
	return nil
}
