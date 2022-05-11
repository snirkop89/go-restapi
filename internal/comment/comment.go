package comment

import (
	"context"
	"fmt"
)

// Comment is a represenation of the comment
// stracture for our service
type Comment struct {
	ID     string
	Slug   string // path to where the comment is
	Body   string
	Author string
}

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
	fmt.Println("Getting comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}

	return cmt, nil
}
