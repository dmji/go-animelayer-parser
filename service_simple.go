package animelayer

import (
	"context"
	"errors"
)

func (p *service) GetItemsFromCategoryPages(ctx context.Context, category Category, iPage int) ([]Item, error) {

	pageNode, err := p.pageTargetToHtmlNode(category, iPage)
	if err != nil {
		return nil, err
	}

	items := make(chan ItemPartialWithError, 100)
	go func() {
		defer close(items)
		p.parserDetailedItems.ParseCategoryPageToChan(ctx, pageNode, items)
	}()

	errs := make([]error, 0, 100)
	res := make([]Item, 0, 100)
	for item := range items {
		if item.Error != nil {
			errs = append(errs, item.Error)
			continue
		}

		res = append(res, *item.Item)
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return res, nil
}

func (p *service) GetItemByIdentifier(ctx context.Context, identifier string) (*Item, error) {
	detailedItemNode := p.partialItemToItemNode(identifier)
	return p.parserDetailedItems.ParseItem(ctx, detailedItemNode.Node, detailedItemNode.Identifier)
}
