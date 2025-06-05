/*
(c) 2025 Tadahiro Ishisaka all rights reserved.

This software is released under the Apache License, Version 2.0.
*/

/*
slogの簡単なサンプルです。
*/
package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	// slogは標準のlogと違いログの出力レベルを指定できるのがメリットの一つ
	// Infoレベルで単純なログを標準エラー出力に出力する
	slog.Info("Hello, world!")
	// プリセットの無いログレベル2のログを出力するにはLog関数を使用する
	slog.Log(context.TODO(), 2, "Hello, world!")

	// slogのメリットの二つ目はログメッセージに属性を含められること
	// ログメッセージに属性を含める
	slog.Info("Hello", "number", 3)
	// slog.Attrを使用する
	// slog.Attrを使用するとリフレクションを使用しなくて済むので
	// 実行速度が速くなる
	slog.Info("hello", slog.Int("number", 3))

	// Logger.Withメソッドを使って新しいLoggerを構築し、全てのレコードにその属性を含める
	logger := slog.Default()
	logger2 := logger.With("url", "https://example.com")
	logger2.Info("Hello, world!")
	// 出力例: 2025/06/03 15:16:13 INFO Hello, world! url=https://example.com

	// JSONハンドラを使ってJSON形式で出力する
	// 標準のハンドラにはJSONとTEXTがある
	loggerwJsonHandler := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	loggerwJsonHandler.Info("Hello, world!")
	// 出力例: {"time":"2025-06-03T15:18:52.53884+09:00","level":"INFO","msg":"Hello, world!"}
	loggerwJsonHandler.Info("Hello", slog.Int("number", 3))
}
