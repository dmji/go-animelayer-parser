package animelayer

import (
	"context"
)

func (p *pipe) PipeItemNodesToDetailedItems(ctx context.Context, itemNodes <-chan ItemNode) <-chan ItemDetailed {
	items := make(chan ItemDetailed, 100)

	go func() {
		defer close(items)
		parser := newParser(p.logger)
		for {

			select {
			case <-ctx.Done():
				return
			case itemNode, bOpen := <-itemNodes:

				if !bOpen && len(itemNodes) == 0 {
					return
				}

				p.logger.Infow("Started item node", "identifier", itemNode.Identifier)
				item := parser.ParseItem(ctx, itemNode.Node)
				if item == nil {
					p.logger.Errorw("Pipe Item Node To Completed Item Error", "error", "got nil item")
				} else {
					items <- *item
				}
			}

		}
	}()

	return items
}
