package animelayer

import (
	"context"
)

func (p *service) PipePartialItemToItemNode(ctx context.Context, partialItems <-chan ItemPartialWithError) <-chan PageHtmlNode {

	itemNodes := make(chan PageHtmlNode, 100)

	go func() {
		defer close(itemNodes)

		for {

			select {
			case <-ctx.Done():
				return
			case prop, bOpen := <-partialItems:

				if !bOpen && len(partialItems) == 0 {
					return
				}

				itemNodes <- p.partialItemToItemNode(*prop.Item)
			}
		}
	}()

	return itemNodes
}
