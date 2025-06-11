package infra

import (
	"context"
	"math/rand"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	t.Run("should work", func(t *testing.T) {
		ctx := t.Context()
		ctx2 := context.WithValue(ctx, "key", "value2")
		ctx3, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		select {
		case <-time.After(time.Duration(rand.Intn(5)) * time.Second):
			t.Logf("%v", ctx2.Value("key"))
			break

		case <-ctx3.Done():
			t.Log("context cancelled")
		}

	})
}
