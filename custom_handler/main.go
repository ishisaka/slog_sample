/*
(c) 2025 Tadahiro Ishisaka all rights reserved.

This software is released under the Apache License, Version 2.0.
*/

/*
複数先へのログ出力することを例にしたカスタムハンドラの例です
*/
package main

import (
	"context"
	"errors" // Go 1.20+ で errors.Join を使う場合
	"log/slog"
	"os"
)

// MultiHandler は複数の slog.Handler にログをディスパッチします。
type MultiHandler struct {
	handlers []slog.Handler
}

// NewMultiHandler は MultiHandler を作成します。
func NewMultiHandler(handlers ...slog.Handler) *MultiHandler {
	return &MultiHandler{handlers: handlers}
}

// Enabled は、いずれかのラップされたハンドラが指定されたレベルで有効な場合に true を返します。
func (h *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, handler := range h.handlers {
		if handler.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

// Handle は、ラップされたすべてのハンドラにログレコードを渡します。
// いずれかのハンドラでエラーが発生した場合、エラーを集約して返します。
func (h *MultiHandler) Handle(ctx context.Context, r slog.Record) error {
	var errs []error
	for _, handler := range h.handlers {
		// 各ハンドラが Enabled かどうかをここで再チェックすることもできますが、
		// 通常は Enabled でフィルタリングされた後に Handle が呼ばれることを期待します。
		if err := handler.Handle(ctx, r); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errors.Join(errs...) // Go 1.20+
		// Go 1.19以前の場合は、最初のエラーを返すか、カスタムエラー型でラップするなどの対応が必要です。
		// return fmt.Errorf("%d errors occurred: %v", len(errs), errs)
	}
	return nil
}

// WithAttrs は、ラップされたすべてのハンドラに属性を追加した新しい MultiHandler を返します。
func (h *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandlers := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		newHandlers[i] = handler.WithAttrs(attrs)
	}
	return NewMultiHandler(newHandlers...)
}

// WithGroup は、ラップされたすべてのハンドラにグループを追加した新しい MultiHandler を返します。
func (h *MultiHandler) WithGroup(name string) slog.Handler {
	newHandlers := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		newHandlers[i] = handler.WithGroup(name)
	}
	return NewMultiHandler(newHandlers...)
}

func main() {
	// ファイルへのハンドラ (JSON形式、DEBUGレベル以上)
	logFile, err := os.OpenFile("app_multi.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("ログファイルを開けませんでした", slog.Any("error", err))
		os.Exit(1)
	}
	defer logFile.Close()
	fileHandler := slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	// 標準出力へのハンドラ (Text形式、INFOレベル以上)
	stdoutHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// 標準出力ではソース情報を表示しない例
			if a.Key == slog.SourceKey {
				return slog.Attr{}
			}
			return a
		},
	})

	// MultiHandler を作成
	multiHandler := NewMultiHandler(stdoutHandler, fileHandler)

	// ロガーを作成
	logger := slog.New(multiHandler)

	// ログ出力
	logger.Debug("このデバッグメッセージはファイルにのみ記録されます。") // stdoutHandlerのレベルはINFOなので表示されない
	logger.Info("この情報メッセージは標準出力とファイルの両方に記録されます。", slog.String("user", "admin"))
	logger.Warn("警告: ディスク容量が少なくなっています。", slog.Int("free_gb", 10))
	logger.Error("エラー発生！", slog.String("component", "API"))

}
