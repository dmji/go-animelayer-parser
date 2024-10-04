package animelayer

import (
	"context"
	"sync"
)

func PipeGenericInterceptor[T any](ctx context.Context, inT <-chan T, capacity int, fn func(p *T)) <-chan T {
	outT := make(chan T, capacity)

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

func PipeGeneric[TIn, TOut any](ctx context.Context, inT <-chan TIn, capacity int, fn func(p *TIn) *TOut) <-chan TOut {
	outT := make(chan TOut, capacity)

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

				tp := fn(&t)
				if tp == nil {
					return
				}

				outT <- *tp
			}

		}
	}()

	return outT
}

func PipeGenericWg[TIn, TOut any](ctx context.Context, inT <-chan TIn, capacity int, fn func(ctx context.Context, p TIn, o chan<- TOut)) <-chan TOut {
	outT := make(chan TOut, capacity)

	go func() {
		defer close(outT)
		wg := &sync.WaitGroup{}

		for {

			select {
			case <-ctx.Done():
				return
			case t, bOpen := <-inT:

				if !bOpen && len(inT) == 0 {
					return
				}

				wg.Add(1)
				go func() {
					defer wg.Done()
					fn(ctx, t, outT)
				}()
			}

		}
	}()

	return outT
}
