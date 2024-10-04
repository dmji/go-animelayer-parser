package animelayer

import (
	"context"
	"fmt"
)

func (p *service) PipeItemNodesToDetailedItems(ctx context.Context, itemNodes <-chan PageHtmlNode) <-chan ItemDetailedWithError {
	items := make(chan ItemDetailedWithError, 100)

	go func() {
		defer close(items)
		for {

			select {
			case <-ctx.Done():
				return
			case itemNode, bOpen := <-itemNodes:

				if !bOpen && len(itemNodes) == 0 {
					return
				}

				item := parseItem(ctx, itemNode.Node)
				if item == nil {
					items <- ItemDetailedWithError{Item: nil, Error: fmt.Errorf("got nil item from identifier='%s'", itemNode.Identifier)}
				} else {
					items <- ItemDetailedWithError{Item: item, Error: nil}
				}
			}

		}
	}()

	return items
}
