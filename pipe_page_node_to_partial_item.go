package animelayer

import (
	"context"
	"sync"
)

func (p *service) PipePageNodesToPartialItems(ctx context.Context, pageNodes <-chan CategoryHtml) <-chan ItemPartialWithError {
	items := make(chan ItemPartialWithError, 100)

	go func() {
		wg := &sync.WaitGroup{}
		defer close(items)

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
					parseCategoryPageToChan(ctx, pageNode.Node, items)
				}()
			}
		}
		wg.Wait()
	}()

	return items
}
