# 概要

- gRPC について学習したことを管理するためのリポジトリ

## 環境構築

### Go インストール

- [手順](./install_go.md)

### Go 言語の実装に便利なモジュールのインストール

- 拡張機能 `Go` を入れている場合は、`F1` から `Go: Install/Update tools` を選択して、表示される全てのモジュールをインストールする

  - 以下のコマンドを代わりに実行してくれる感じ

    ```bash
    go install github.com/cweill/gotests/gotests@v1.6.0
    go install github.com/fatih/gomodifytags@v1.16.0
    go install github.com/josharian/impl@v1.1.0
    go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0
    go install github.com/go-delve/delve/cmd/dlv@latest
    go install honnef.co/go/tools/cmd/staticcheck@latest
    go install golang.org/x/tools/gopls@latest
    ```

### gRPC の実装に必要なモジュールのインストール

```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Tips

- Go 言語はコンパイル言語である。そのため、Python における virtualenv 等が不要である。上記モジュールインストール時にも分かるが、モジュール毎に必要なサブモジュールのバイナリがダウンロードされるのが分かる
  - FYI: https://eli.thegreenplace.net/2020/you-dont-need-virtualenv-in-go/

## protocol buffers の概要

- [README.md](./protobuf-lesson/README.md)

## gRPC の概要

- [README.md](./grpc-lesson/README.md)
