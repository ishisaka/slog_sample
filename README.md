# slog_sample

Go言語のlog/slogサンプル

## log/slogについて

Go言語の`log/slog`パッケージは、Go 1.21で導入された**構造化ロギング**のための標準ライブラリです。ログをキーと値のペアとして出力することで、人間にとっての可読性だけでなく、機械による処理や分析も容易にします。

---
### 主な特徴 📝

* **レベル付きロギング**: `Debug`, `Info`, `Warn`, `Error`といったログレベルをサポートし、出力するログの重要度を制御できます。
* **構造化された出力**: デフォルトでは、ログはキーと値のペアで構成されるテキスト形式（`key=value`）またはJSON形式で出力されます。これにより、ログの検索やフィルタリングが容易になります。
* **柔軟なハンドラ**: `slog.Handler`インターフェースを実装することで、ログの出力形式（テキスト、JSONなど）や出力先（標準出力、ファイル、ネットワークなど）をカスタマイズできます。
* **コンテキスト対応**: `context.Context`と連携し、リクエストIDなどの共通情報をログに含めることが容易です。

---
### 簡単な使い方 💡

```go
package main

import (
	"log/slog"
	"os"
)

func main() {
	// デフォルトのテキストハンドラでロガーを作成
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// 情報レベルでログを出力
	logger.Info("ユーザーがログインしました", "userID", 123, "userName", "gopher")

	// エラーレベルでログを出力
	logger.Error("データベース接続に失敗しました", "error", "connection refused")
}
```

---
### メリット ✨

`log/slog`を利用することで、ログの可読性向上、効率的なログ分析、パフォーマンスへの影響低減といったメリットが期待できます。標準パッケージであるため、外部ライブラリへの依存を減らすことも可能です。

より詳細な情報や高度な使い方については、[公式ドキュメント](https://pkg.go.dev/log/slog)を参照してください。

## ディレクトリ構成

- simple_sample
  - slogのシンプルな使い方
- config_file:
  - slogの設定に設定ファイルを使用する例 

## 参考

- [slog package \- golang\.org/x/exp/slog \- Go Packages](https://pkg.go.dev/golang.org/x/exp/slog)
