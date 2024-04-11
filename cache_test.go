package ctxcache

import (
	"context"
	"sync"
	"testing"
)

func TestCache(t *testing.T) {
	ctx := context.Background()
	ctx = NewContextWithCache(ctx)
	val, ok := Get(ctx, "isBelongTo")
	t.Log(val, ok)
	_ = Set(ctx, "isBelongTo", true)
	_ = Set(ctx, "lawlielt", "123")
	val, ok = Get(ctx, "isBelongTo")
	t.Log(val, ok)
}

func TestCacheConcurrent(t *testing.T) {
	ctx := context.Background()
	ctx = NewContextWithCache(ctx)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val, ok := Get(ctx, "isBelongTo")
			if ok {
				t.Log(val)
				return
			}
			_ = Set(ctx, "isBelongTo", i)
		}(i)
	}
	wg.Wait()
}
