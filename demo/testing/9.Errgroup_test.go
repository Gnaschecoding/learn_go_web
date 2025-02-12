package testing

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync/atomic"
	"testing"
)

func TestErrgroup(t *testing.T) {
	//eg := errgroup.Group{}
	eg, ctx := errgroup.WithContext(context.Background())
	var result int64 = 0
	for i := 0; i < 10; i++ {
		delta := 1
		eg.Go(func() error {
			atomic.AddInt64(&result, int64(delta))
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
	ctx.Err()
	fmt.Println(result)
}
