package animelayer

import (
	"context"
	"fmt"
)

func (p *service) PipeItemNodesToDetailedItems(ctx context.Context, itemNodes <-chan PageHtmlNode) <-chan ItemDetailedWithError {

	return PipeGeneric(ctx, itemNodes, 100, func(itemNode *PageHtmlNode) *ItemDetailedWithError {

		item := parseItem(ctx, itemNode.Node)
		if item == nil {
			return &ItemDetailedWithError{Item: nil, Error: fmt.Errorf("got nil item from identifier='%s'", itemNode.Identifier)}
		} else {
			return &ItemDetailedWithError{Item: item, Error: nil}
		}

	})

}
