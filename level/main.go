/*
(c) 2025 Tadahiro Ishisaka all rights reserved.

This software is released under the Apache License, Version 2.0.
*/

/*
slogのレベル設定のサンプル
*/
package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	// ログレベルの変数を作る
	var programLevel = new(slog.LevelVar)
	// ログハンドラーを作り、ログレベルの変数をHandlerOptionのLevelに割り当てる
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	// 作成したハンドラーを標準のログ出力に設定
	slog.SetDefault(slog.New(h))

	/*
		Log Levelの設定
		const (
			LevelDebug Level = -4
			LevelInfo  Level = 0
			LevelWarn  Level = 4
			LevelError Level = 8
		)
	*/
	// ログレベルを設定
	programLevel.Set(slog.LevelDebug)

	// 設定したレベル以上のログが表示される
	slog.Debug("debug")
	slog.Info("info")
	slog.Warn("warn")
	slog.Error("error")

	// 標準には無いログレベルの設定と出力
	slog.Log(context.TODO(), 2, "Level 2")
	// 表示例： {"time":"2025-06-04T10:04:03.603603+09:00","level":"INFO+2","msg":"Level 2"}
}
