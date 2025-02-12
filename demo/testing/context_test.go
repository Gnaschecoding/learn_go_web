package testing

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctx := context.Background()

	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	time.Sleep(time.Second * 2)
	dl, err := timeoutCtx.Deadline()
	fmt.Println(dl, err)
}
