/*
(c) 2025 Tadahiro Ishisaka all rights reserved.

This software is released under the Apache License, Version 2.0.
*/

/*
ログレコードの属性を置き換えるサンプル
https://pkg.go.dev/golang.org/x/exp/slog#HandlerOptions
*/
package main

import (
	"os"

	"log/slog"
)

func main() {
	// Exported constants from a custom logging package.
	const (
		LevelTrace     = slog.Level(-8)
		LevelDebug     = slog.LevelDebug
		LevelInfo      = slog.LevelInfo
		LevelNotice    = slog.Level(2)
		LevelWarning   = slog.LevelWarn
		LevelError     = slog.LevelError
		LevelEmergency = slog.Level(12)
	)

	th := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		// カスタムレベルのLevelTraceをデフォルトのレベルに設定
		Level: LevelTrace,

		// ログレコードの属性を置き換える関数を設定
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// ログ出力からタイムスタンプを削除
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}

			// カスタムのレベルとそのキーを設定
			if a.Key == slog.LevelKey {
				// レベルのキーを"level" から "sev" に変更
				a.Key = "sev"

				// カスタムレベル値のハンドする
				level := a.Value.Any().(slog.Level)

				// カスタムレベルに合わせてログに出力するレベルを表す文字列を設定
				switch {
				case level < LevelDebug:
					a.Value = slog.StringValue("TRACE")
				case level < LevelInfo:
					a.Value = slog.StringValue("DEBUG")
				case level < LevelNotice:
					a.Value = slog.StringValue("INFO")
				case level < LevelWarning:
					a.Value = slog.StringValue("NOTICE")
				case level < LevelError:
					a.Value = slog.StringValue("WARNING")
				case level < LevelEmergency:
					a.Value = slog.StringValue("ERROR")
				default:
					a.Value = slog.StringValue("EMERGENCY")
				}
			}

			return a
		},
	})

	logger := slog.New(th)
	logger.Log(nil, LevelEmergency, "missing pilots")
	logger.Error("failed to start engines", "err", "missing fuel")
	logger.Warn("falling back to default value")
	logger.Log(nil, LevelNotice, "all systems are running")
	logger.Info("initiating launch")
	logger.Debug("starting background job")
	logger.Log(nil, LevelTrace, "button clicked")

}

/*
出力例
sev=EMERGENCY msg="missing pilots"
sev=ERROR msg="failed to start engines" err="missing fuel"
sev=WARNING msg="falling back to default value"
sev=NOTICE msg="all systems are running"
sev=INFO msg="initiating launch"
sev=DEBUG msg="starting background job"
sev=TRACE msg="button clicked"
*/
