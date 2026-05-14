// Package context provides cancellation and timeout support for logslice
// pipeline operations, allowing long-running file processing to be interrupted
// gracefully via signal handling or explicit cancellation.
package context

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// WithSignals returns a context that is cancelled when the process receives
// SIGINT or SIGTERM. The returned stop function must be called to release
// resources associated with the signal handler.
func WithSignals(parent context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		select {
		case <-ch:
			cancel()
		case <-ctx.Done():
		}
		signal.Stop(ch)
	}()

	stop := func() {
		cancel()
		signal.Stop(ch)
	}

	return ctx, stop
}

// IsCancelled reports whether the given context has been cancelled or has
// exceeded its deadline.
func IsCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
