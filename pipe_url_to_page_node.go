package animelayer

import (
	"context"
)

func (p *service) PipePagesFromCategoryToPageNode(ctx context.Context, category Category, pages ...int) <-chan CategoryHtml {
	documents := make(chan CategoryHtml, 10)

	getPageIndex := func(i int) (int, bool) {
		if len(pages) == 0 {
			return i, false
		}

		if len(pages) <= i {
			return 0, true
		}

		return pages[i], false
	}

	go func() {
		defer close(documents)

		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			default:
			}

			iPage, bBreak := getPageIndex(i)
			if bBreak {
				break
			}

			doc, err := p.pageTargetToHtmlNode(category, iPage)
			if err != nil {
				break
			}

			documents <- *doc
		}

	}()

	return documents
}
