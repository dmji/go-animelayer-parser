package animelayer

import (
	"context"
)

func (p *service) partialItemToItemNode(item ItemPartial) ItemNode {
	url := formatUrlToItem(item.Identifier)

	p.logger.Infow("Started item partial", "url", url)
	doc, err := loadHtmlDocument(p.client, url)
	if err != nil {
		panic(err)
	}
	return ItemNode{Node: doc, Identifier: item.Identifier}
}

func (p *service) PipePartialItemToItemNode(ctx context.Context, partialItems <-chan ItemPartial) <-chan ItemNode {

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

				itemNodes <- p.partialItemToItemNode(item)
			}
		}
	}()

	return itemNodes
}
