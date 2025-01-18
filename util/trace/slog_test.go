package trace

import (
	"log/slog"
	"testing"
)

func TestContextHandler_Handle(t *testing.T) {
	slog.Info("你好", "你好", "你好")

	ctx := Context()
	slog.WarnContext(ctx, "测试 warn", "123", "value 123")
}