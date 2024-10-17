package animelayer

import (
	"context"
)

type Parser interface {
	GetItemByIdentifier(ctx context.Context, identifier string) (*Item, error)
	GetItemsFromCategoryPages(ctx context.Context, category Category, iPage int) ([]Item, error)
}

type service struct {
	client              HtmlDocGetter
	parserDetailedItems parserDetailedItems
}

func New(client HtmlDocGetter, ServiceOptionsApplier ...ServiceOptionsApplier) *service {
	s := &service{
		client: client,
	}

	for _, apply := range ServiceOptionsApplier {
		apply(s)
	}

	return s
}
