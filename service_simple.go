package animelayer

import (
	"context"
	"errors"
)

func (p *service) CategoryPageToPartialItems(ctx context.Context, category Category, iPage int) ([]ItemPartial, error) {

	pageNode, err := p.pageTargetToHtmlNode(category, iPage)
	if err != nil {
		return nil, err
	}

	items := make(chan ItemPartialWithError, 100)
	go func() {
		defer close(items)
		parseCategoryPageToChan(ctx, pageNode, items)
	}()

	errs := make([]error, 0, 100)
	res := make([]ItemPartial, 0, 100)
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

func (p *service) PartialItemToDetailedItem(ctx context.Context, identifier string) (*ItemDetailed, error) {
	detailedItemNode := p.partialItemToItemNode(identifier)
	return p.parserDetailedItems.parseItem(ctx, detailedItemNode.Node, detailedItemNode.Identifier)
}
