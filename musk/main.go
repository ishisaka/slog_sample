/*
(c) 2025 Tadahiro Ishisaka all rights reserved.

This software is released under the Apache License, Version 2.0.
*/

/*
トークンなどのシークレットをログ出力時にマスクしたいときの処理
*/

package main

import (
	"log/slog"
	"os"
)

// Token はシークレットにしたいものの例
type Token string

// LogValue はLogValuerインターフェイスの実装
// Token型の値を"REDACTED_TOKEN"に書き換える
func (Token) LogValue() slog.Value {
	return slog.StringValue("REDACTED_TOKEN")
}

func main() {
	t := Token("shhhh!")
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("permission granted", "user", "Perry", "token", t)
	// time=2025-06-04T14:06:36.595+09:00 level=INFO msg="permission granted" user=Perry token=REDACTED_TOKEN
}
