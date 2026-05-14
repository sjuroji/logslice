// Package context wraps the standard library context package to provide
// logslice-specific helpers for graceful cancellation.
//
// # Signal handling
//
// WithSignals creates a context that is automatically cancelled when the
// process receives SIGINT or SIGTERM, enabling the pipeline to stop cleanly
// without leaving partial output:
//
//	ctx, stop := logctx.WithSignals(context.Background())
//	defer stop()
//
//	if err := pipeline.Run(ctx, cfg); err != nil {
//		// err may be context.Canceled on clean shutdown
//	}
//
// # Cancellation check
//
// IsCancelled provides a non-blocking test suitable for use inside tight
// processing loops where checking ctx.Done() via a select would be verbose.
package context
