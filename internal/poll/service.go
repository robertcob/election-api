package poll

import (
	"context"
	"election-api/internal/entity"
)

type Service interface {
	Get(ctx context.Context, id string) (Poll, error)
	Query(ctx context.Context, offset, limit int) ([]Poll, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreatePollRequest) (Poll, error)
	Update(ctx context.Context, id string, input UpdatePollRequest) (Poll, error)
	Delete(ctx context.Context, id string) (Poll, error)
}

type Poll struct {
	entity.Poll
}

type CreatePollRequest struct {
}

type UpdatePollRequest struct {
}
