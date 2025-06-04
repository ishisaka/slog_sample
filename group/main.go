/*
(c) 2025 Tadahiro Ishisaka all rights reserved.

This software is released under the Apache License, Version 2.0.
*/

/*
slogのGroupを使う場合のサンプル
Groupはログメッセージに属性を含める際に便利です。
*/
package main

import (
	"os"

	"log/slog"
)

func main() {
	// テキストハンドラーを作成（標準エラー出力）
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	// JSONハンドラーを作成したい場合は以下のようにします
	// logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))

	// --- グループを使用したログ出力の例 ---
	logger.Info("ユーザー情報",
		slog.String("id", "user-123"),
		slog.Group("details",
			slog.String("email", "test@example.com"),
			slog.Int("age", 30),
			slog.Group("address", // ネストしたグループ
				slog.String("street", "123 Main St"),
				slog.String("city", "Anytown"),
			),
		),
		slog.Bool("isActive", true),
	)

	// --- WithGroup を使用してロガーにグループを永続的に設定する例 ---
	// "request" グループを持つ新しいロガーを作成
	requestLogger := logger.WithGroup("request")

	requestLogger.Info("受信リクエスト",
		slog.String("method", "GET"),
		slog.String("path", "/api/data"),
	)

	requestLogger.Error("リクエスト処理エラー",
		slog.Int("statusCode", 500),
		slog.String("error", "データベース接続エラー"),
	)

	// さらにネストしたグループを WithGroup で設定
	userRequestLogger := requestLogger.WithGroup("user")
	userRequestLogger.Info("ユーザー関連リクエスト",
		slog.String("userID", "user-456"),
	)

}
