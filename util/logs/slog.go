package logs

import (
	"context"
	"log/slog"
	"os"

	"sbp/util/idh"
)

type ContextHandler struct {
	slog.Handler
}

func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	// context 需要在首次出现地方 注入 idh.WithTraceID(&ctx) trace id
	traceID := idh.GetTraceID(ctx)
	if len(traceID) > 0 {
		r.AddAttrs(slog.String(idh.TraceIDKey, traceID))
	}

	return h.Handler.Handle(ctx, r)
}

func init() {
	opts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logs := slog.New(&ContextHandler{slog.NewJSONHandler(os.Stdout, &opts)})
	slog.SetDefault(logs)
}