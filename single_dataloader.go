package dataloader

import (
	"context"
)

type SingleLoad[K comparable, V any] interface {
	Load(context.Context, K) Thunk[V]
	Clear(context.Context, K) Interface[K, V]
	ClearAll() Interface[K, V]
	Prime(ctx context.Context, key K, value V) Interface[K, V]
}

type FetchFunc[K comparable, V any] func(context.Context, K) *Result[V]

func NewSingleLoader[K comparable, V any](loadFn FetchFunc[K, V], opts ...Option[K, V]) SingleLoad[K, V] {
	batchFunc := func(ctx context.Context, keys []K) []*Result[V] {
		var result []*Result[V]
		for _, key := range keys {
			result = append(result, loadFn(ctx, key))
		}
		return nil
	}
	return NewBatchedLoader(batchFunc, opts...)
}
