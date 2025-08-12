package util

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
)

// WithNotifyContext wraps signal.NotifyContext. It returns a copy of the parent
// context that is canceled when the process receives a SIGTERM, os.Interrupt, or
// os.Kill signal.
func WithNotifyContext(parent context.Context) (ctx context.Context, stop context.CancelFunc) {
	return signal.NotifyContext(parent, syscall.SIGTERM, os.Interrupt, os.Kill)
}

var ErrChClosed = errors.New("ch closed")

// RecvContext receives a value from the channel ch.
// It returns when a value is received or the context is done.
// If the context is done, it returns the context's error.
// If the channel is closed while waiting for a value, it returns ErrChClosed.
func RecvContext[T any](ctx context.Context, ch <-chan T) (T, error) {
	var zero T
	select {
	case <-ctx.Done():
		return zero, ctx.Err()
	case v, ok := <-ch:
		if !ok {
			return zero, ErrChClosed
		}
		return v, nil
	}
}
