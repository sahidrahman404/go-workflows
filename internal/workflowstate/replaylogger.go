package workflowstate

import (
	"context"
	"log/slog"
)

type replayHandler struct {
	state   *WfState
	handler slog.Handler
}

// Enabled implements slog.Handler.
func (rh *replayHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return rh.handler.Enabled(ctx, level)
}

// Handle implements slog.Handler.
func (rh *replayHandler) Handle(ctx context.Context, r slog.Record) error {
	if rh.state.Replaying() {
		return nil
	}

	return rh.handler.Handle(ctx, r)
}

// WithAttrs implements slog.Handler.
func (rh *replayHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return rh.handler.WithAttrs(attrs)
}

// WithGroup implements slog.Handler.
func (rh *replayHandler) WithGroup(name string) slog.Handler {
	return rh.handler.WithGroup(name)
}

var _ slog.Handler = (*replayHandler)(nil)

func NewReplayLogger(state *WfState, logger *slog.Logger) *slog.Logger {
	h := logger.Handler()

	return slog.New(&replayHandler{state, h})
}
