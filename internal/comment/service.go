package comment

import "context"

// Service - a service for comments
type Service struct {
	repository Repository
}

// NewService - returns a new instance of a comment service
func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// Get - returns a comment by id
func (s *Service) Get(ctx context.Context, id string) (Comment, error) {
	result, err := s.repository.Get(ctx, id)
	if err != nil {
		return Comment{}, err
	}
	comment := Comment{
		ID:     result.ID,
		Slug:   result.Slug,
		Body:   result.Body,
		Author: result.Author,
	}
	return comment, nil
}

// Delete - deletes a comment by id
func (s *Service) Delete(ctx context.Context, id string) (Comment, error) {
	result, err := s.repository.Delete(ctx, id)
	if err != nil {
		return Comment{}, err
	}
	comment := Comment{
		ID:     result.ID,
		Slug:   result.Slug,
		Body:   result.Body,
		Author: result.Author,
	}
	return comment, nil
}

// Create - creates a new comment
func (s *Service) Create(ctx context.Context, c Comment) (Comment, error) {
	result, err := s.repository.Create(ctx, c)
	if err != nil {
		return Comment{}, err
	}
	comment := Comment{
		ID:     result.ID,
		Slug:   result.Slug,
		Body:   result.Body,
		Author: result.Author,
	}
	return comment, nil
}

// Update - updates a comment by id
func (s *Service) Update(ctx context.Context, id string, c Comment) (Comment, error) {
	result, err := s.repository.Update(ctx, id, c)
	if err != nil {
		return Comment{}, err
	}
	comment := Comment{
		ID:     result.ID,
		Slug:   result.Slug,
		Body:   result.Body,
		Author: result.Author,
	}
	return comment, nil
}

// List - returns all comments
func (s *Service) List(ctx context.Context) (Comments, error) {
	result, err := s.repository.List(ctx)
	if err != nil {
		return Comments{}, err
	}
	comments := Comments{}
	for _, record := range result {
		comments = append(comments, Comment{
			ID:     record.ID,
			Slug:   record.Slug,
			Body:   record.Body,
			Author: record.Author,
		})
	}
	return comments, nil
}
