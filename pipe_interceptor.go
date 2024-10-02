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
