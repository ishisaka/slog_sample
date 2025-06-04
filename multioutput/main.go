/*
(c) 2025 Tadahiro Ishisaka all rights reserved.

This software is released under the Apache License, Version 2.0.
*/

/*
ログを複数の出力先に出力するサンプルです。
*/
package main

import (
	"io"
	"log/slog"
	"os"
)

func main() {
	// ログをファイルと標準出力に同時に出力する
	// ログファイルを開く (または作成する)
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("ログファイルを開けませんでした", slog.Any("error", err))
		os.Exit(1)
	}
	defer logFile.Close()

	// io.MultiWriter を作成して、標準出力とファイルの両方に出力するようにする
	//    os.Stdout: 標準出力
	//    logFile:   開いたファイル
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// MultiWriter を出力先とするハンドラを作成
	//    ここでは TextHandler を使用する例。JSONHandler も同様に使えます。
	//    handler := slog.NewJSONHandler(multiWriter, nil)
	handler := slog.NewTextHandler(multiWriter, &slog.HandlerOptions{
		AddSource: true,            // ソースコードの位置情報を追加する場合
		Level:     slog.LevelDebug, // ログレベルを設定
	})

	// カスタムハンドラを使用してロガーを作成
	logger := slog.New(handler)

	// ログを出力 (標準出力と app.log の両方に出力される)
	logger.Debug("これはデバッグメッセージです。")
	logger.Info("アプリケーションが起動しました。", slog.String("version", "1.0.0"))
	logger.Warn("設定ファイルが見つかりません。デフォルト値を使用します。", slog.String("config_path", "./config.toml"))
	logger.Error("重大なエラーが発生しました。", slog.String("error_code", "SYS_001"), slog.String("details", "データベース接続に失敗"))

	slog.SetDefault(logger)
	slog.Info("デフォルトロガーも設定できます。") // slog.SetDefault で設定した場合
}
