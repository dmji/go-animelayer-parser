package animelayer

import (
	"context"
)

// ItemProvider - interface of main package object
type ItemProvider interface {
	GetItemByIdentifier(ctx context.Context, identifier string) (*Item, error)
	GetItemsFromCategoryPages(ctx context.Context, category Category, iPage int) ([]Item, error)
}

type service struct {
	client DocGetter
	parser parser
}

// New - create main package object
func New(client DocGetter, ServiceOptionsApplier ...ServiceOptionsApplier) ItemProvider {
	s := &service{
		client: client,
	}

	for _, apply := range ServiceOptionsApplier {
		apply(s)
	}

	return s
}
