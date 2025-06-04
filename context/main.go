/*
(c) 2025 Tadahiro Ishisaka all rights reserved.

This software is released under the Apache License, Version 2.0.
*/

/*
ログにcontextの内容を追加するサンプルです。
*/
package main

import (
	"context"
	"log/slog"
	"os"
)

// contextKey はコンテキスト内で値を一意に識別するためのキーです。
type contextKey string

const (
	traceIDKey contextKey = "traceID"
	userIDKey  contextKey = "userID"
)

// ContextHandler は、コンテキストから指定されたキーの値を抽出し、
// ログレコードに自動的に追加する slog.Handler のラッパーです。
type ContextHandler struct {
	slog.Handler
	keys []contextKey // コンテキストから抽出するキーのリスト
}

// NewContextHandler は ContextHandler を作成します。
func NewContextHandler(handler slog.Handler, keys []contextKey) *ContextHandler {
	return &ContextHandler{
		Handler: handler,
		keys:    keys,
	}
}

// Handle は、元のハンドラの Handle メソッドを呼び出す前に、
// コンテキストから情報を抽出してログレコードに追加します。
func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, key := range h.keys {
		if val := ctx.Value(key); val != nil {
			// slog.Any は便利ですが、具体的な型で slog.String, slog.Int などを使う方が望ましい場合もあります。
			r.AddAttrs(slog.Any(string(key), val))
		}
	}
	return h.Handler.Handle(ctx, r)
}

// WithAttrs は、ラップされたハンドラの WithAttrs を呼び出します。
// 新しい ContextHandler を返すように実装することもできます。
func (h *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewContextHandler(h.Handler.WithAttrs(attrs), h.keys)
}

// WithGroup は、ラップされたハンドラの WithGroup を呼び出します。
// 新しい ContextHandler を返すように実装することもできます。
func (h *ContextHandler) WithGroup(name string) slog.Handler {
	return NewContextHandler(h.Handler.WithGroup(name), h.keys)
}

func main() {
	// ベースとなるハンドラ (例: JSONHandler)
	baseHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true, // ソースコードの位置情報を追加
		Level:     slog.LevelDebug,
	})

	// コンテキストから traceIDKey と userIDKey を抽出するカスタムハンドラ
	contextAwareHandler := NewContextHandler(baseHandler, []contextKey{traceIDKey, userIDKey})

	// カスタムハンドラを使用してロガーを作成
	logger := slog.New(contextAwareHandler)

	// --- コンテキストに情報を追加 ---
	ctx := context.Background()
	ctx = context.WithValue(ctx, traceIDKey, "trace-xyz-789")
	ctx = context.WithValue(ctx, userIDKey, "user-prod-456")

	// --- ログ出力 (カスタムハンドラが自動的にコンテキスト情報を追加) ---
	logger.InfoContext(ctx, "ユーザーがログインしました", slog.String("username", "gopher"))

	// 別のリクエストのコンテキスト（異なる値）
	ctx2 := context.Background()
	ctx2 = context.WithValue(ctx2, traceIDKey, "trace-def-456")
	// userIDKey はこのコンテキストには設定しない

	logger.WarnContext(ctx2, "在庫が少なくなっています", slog.String("itemID", "item-001"), slog.Int("currentStock", 5))

	// コンテキストがない場合（またはキーが含まれていない場合）
	logger.ErrorContext(context.Background(), "重要なエラーが発生しました", slog.String("errorCode", "E-1024"))
}
