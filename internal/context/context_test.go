package context_test

import (
	"context"
	"testing"
	"time"

	logctx "github.com/yourorg/logslice/internal/context"
)

func TestWithSignals_CancelViaStop(t *testing.T) {
	ctx, stop := logctx.WithSignals(context.Background())
	defer stop()

	if logctx.IsCancelled(ctx) {
		t.Fatal("context should not be cancelled before stop")
	}

	stop()

	select {
	case <-ctx.Done():
		// expected
	case <-time.After(500 * time.Millisecond):
		t.Fatal("context was not cancelled after stop()")
	}
}

func TestWithSignals_ParentCancellation(t *testing.T) {
	parent, parentCancel := context.WithCancel(context.Background())
	ctx, stop := logctx.WithSignals(parent)
	defer stop()

	parentCancel()

	select {
	case <-ctx.Done():
		// expected: child follows parent
	case <-time.After(500 * time.Millisecond):
		t.Fatal("child context was not cancelled when parent was cancelled")
	}
}

func TestIsCancelled_NotCancelled(t *testing.T) {
	ctx := context.Background()
	if logctx.IsCancelled(ctx) {
		t.Fatal("Background context should not appear cancelled")
	}
}

func TestIsCancelled_Cancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if !logctx.IsCancelled(ctx) {
		t.Fatal("cancelled context should be reported as cancelled")
	}
}

func TestIsCancelled_DeadlineExceeded(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()

	time.Sleep(5 * time.Millisecond)

	if !logctx.IsCancelled(ctx) {
		t.Fatal("timed-out context should be reported as cancelled")
	}
}
