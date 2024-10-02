package animelayer

import (
	"context"
)

func (p *pipe) PipePartialItemToItemNode(ctx context.Context, partialItems <-chan ItemPartial) <-chan ItemNode {

	itemNodes := make(chan ItemNode, 100)

	go func() {
		defer close(itemNodes)

		for {

			select {
			case <-ctx.Done():
				return
			case item, bOpen := <-partialItems:

				if !bOpen && len(partialItems) == 0 {
					return
				}

				url := formatUrlToItem(item.Identifier)

				p.logger.Infow("Started item partial", "url", url)
				doc, err := loadHtmlDocument(p.client, url)
				if err != nil {
					panic(err)
				}

				itemNodes <- ItemNode{Node: doc, Identifier: item.Identifier}
			}
		}
	}()

	return itemNodes
}
