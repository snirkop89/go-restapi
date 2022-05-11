package comment

import (
	"context"
	"errors"
	"log"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment is a represenation of the comment
// stracture for our service
type Comment struct {
	ID     string
	Slug   string // path to where the comment is
	Body   string
	Author string
}

// Store defines all the methods which our
// service needs to implement
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service is the struct which on all our
// comment logic is implemented on
type Service struct {
	Store
}

// NewService returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		log.Println(err)
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmd Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
