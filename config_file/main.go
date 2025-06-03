/*
(c) 2025 Tadahiro Ishisaka all rights reserved.

This software is released under the Apache License, Version 2.0.
*/

/*
設定ファイルを使用してslogのログレベルやログ出力形式を設定するサンプルです。
以下のようなJSON形式のファイルを使用します。

ファイル名: config.json
{
  "log_level": "debug",
  "log_format": "json",
  "add_source": true
}
*/

package main

import (
	"encoding/json"
	"log/slog"
	"os"
)

type LogConfig struct {
	LogLevel  string `json:"log_level"`
	LogFormat string `json:"log_format"`
	AddSource bool   `json:"add_source"`
}

func main() {
	// 設定ファイルの読み込み
	configFile, err := os.Open("config.json")
	if err != nil {
		slog.Error("Failed to open config file", "error", err)
		os.Exit(1)
	}
	defer configFile.Close()

	var config LogConfig
	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		slog.Error("Failed to decode config file", "error", err)
		os.Exit(1)
	}

	// ログレベルの設定
	var level slog.Level
	switch config.LogLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		slog.Warn("Invalid log level in config, defaulting to Info", "configured_level", config.LogLevel)
		level = slog.LevelInfo
	}

	// ログ出力形式の設定
	opts := &slog.HandlerOptions{
		// 出力元のファイル名と行番号を追加するかどうかの設定
		AddSource: config.AddSource,
		Level:     level,
	}

	// ログ出力形式の設定に応じてロガーの作成
	var handler slog.Handler
	switch config.LogFormat {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	case "text":
		handler = slog.NewTextHandler(os.Stdout, opts)
	default:
		slog.Warn("Invalid log format in config, defaulting to Text", "configured_format", config.LogFormat)
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	// 新しいロガーを作成しそれ標準のロガーにする
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// 設定されたロガーでログ出力
	slog.Debug("This is a debug message.")
	slog.Info("This is an info message.")
	slog.Warn("This is a warning message.")
	slog.Error("This is an error message.")
}
