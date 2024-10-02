package animelayer

import (
	"context"
)

func PipeGenericInterceptor[T any](ctx context.Context, inT <-chan T, fn func(p *T)) <-chan T {
	outT := make(chan T, 100)

	go func() {
		defer close(outT)

		for {

			select {
			case <-ctx.Done():
				return
			case t, bOpen := <-inT:

				if !bOpen && len(inT) == 0 {
					return
				}

				fn(&t)

				outT <- t
			}
		}
	}()

	return outT
}

/*
	var b bytes.Buffer
		err := html.Render(&b, pageNode.Node)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(fmt.Sprintf("category_anime/page_%.3d.html", pageNode.Page), b.Bytes(), 0644)
		if err != nil {
			panic(err)
		}
*/

/*
	var b bytes.Buffer
	err := html.Render(&b, itemNode.Node)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fmt.Sprintf("category_anime/page_%.3d.html", itemNode.Identifier), b.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
*/
