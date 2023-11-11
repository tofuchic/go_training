# 概要

- gRPC について学習したことを管理するためのリポジトリ

## 環境構築

### Go インストール

- [手順](./install_go.md)

### gRPC の実装に必要なモジュールのインストール

```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Tips

- Go 言語はコンパイル言語である。そのため、Python における virtualenv 等が不要である。上記モジュールインストール時にも分かるが、モジュール毎に必要なサブモジュールのバイナリがダウンロードされるのが分かる
  - FYI: https://eli.thegreenplace.net/2020/you-dont-need-virtualenv-in-go/
