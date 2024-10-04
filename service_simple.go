package animelayer

import "context"

func (p *service) CategoryPageToPartialItems(ctx context.Context, category Category, iPage int) ([]ItemPartialWithError, error) {

	pageNode, err := p.pageTargetToHtmlNode(category, iPage)
	if err != nil {
		return nil, err
	}

	items := make(chan ItemPartialWithError, 100)
	go func() {
		defer close(items)
		parseCategoryPageToChan(ctx, pageNode, items)
	}()

	res := make([]ItemPartialWithError, 0, 100)
	for item := range items {
		res = append(res, item)
	}

	return res, nil
}

func (p *service) PartialItemToDetailedItem(ctx context.Context, identifier string) *ItemDetailed {
	detailedItemNode := p.partialItemToItemNode(identifier)
	return parseItem(ctx, detailedItemNode.Node)
}
