package animelayer

import "context"

func (p *service) CategoryPageToPartialItems(ctx context.Context, category category, iPage int) ([]ItemPartial, error) {

	pageNode, err := p.pageTargetToHtmlNode(category, iPage)
	if err != nil {
		return nil, err
	}

	parser := newParser(p.logger)

	items := make(chan ItemPartial, 100)
	go func() {
		defer close(items)
		parser.ParsePartialItemsToChan(ctx, pageNode.Node, items)
	}()

	res := make([]ItemPartial, 0, 100)
	for item := range items {
		res = append(res, item)
	}

	return res, nil
}

func (p *service) PartialItemToDetailedItem(ctx context.Context, partialItem ItemPartial) *ItemDetailed {
	parser := newParser(p.logger)

	detailedItemNode := p.partialItemToItemNode(partialItem)
	return parser.ParseItem(ctx, detailedItemNode.Node)
}
