package comment

import (
	"context"
	"errors"
)

var (
	ErrorGetComment = errors.New("could not get comment")
)

// Comment - a representation of a comment
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - defines all of the methods that our service needs
type Store interface {
	GetComment(ctx context.Context, id string) (Comment, error)
}

// Service - a service for comments
type Service struct {
	Store Store
}

// NewService - returns a new instance of a comment service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	comment, err := s.Store.GetComment(ctx, id)
	if err != nil {
		return Comment{}, ErrorGetComment
	}
	return comment, nil
}
