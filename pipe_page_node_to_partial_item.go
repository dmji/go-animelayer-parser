package animelayer

import (
	"context"
	"sync"
)

func (p *service) PipePageNodesToPartialItems(ctx context.Context, pageNodes <-chan PageNode) <-chan ItemPartial {
	items := make(chan ItemPartial, 100)

	go func() {
		wg := &sync.WaitGroup{}
		defer close(items)
		parser := newParser(p.logger)

	loop:
		for {

			select {
			case <-ctx.Done():
				break loop
			case pageNode, bOpen := <-pageNodes:

				if !bOpen && len(pageNodes) == 0 {
					break loop
				}

				wg.Add(1)
				go func() {
					defer wg.Done()
					p.logger.Infow("Started page node", "page", pageNode.Page)
					parser.ParsePartialItemsToChan(ctx, pageNode.Node, items)
				}()
			}
		}
		wg.Wait()
	}()

	return items
}
