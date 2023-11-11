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
