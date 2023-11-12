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


## gRPCの概要

- Google製のRPCで、OSS
  - gはGoogleのgではない
- 型定義にProtocol Buffersを使用していて、各言語毎の型定義のソースコードを自動生成できる
  - 言語間で違いを感じることなく
- HTTP/2を使っているため高速
- 開発の流れ
  1. protoファイルを作成
  1. protoファイルをコンパイルして、クライアント・サーバの雛型となるソースコード（型定義）を生成
  1. 生成された型定義を元にServiceを作成
    - Serviceに1つ以上のメソッドを実装する
  1. Serviceをコンパイルしてインタフェースを作成

### 4つの通信方式とServiceの記述例

- Unary RPC

  ```mermaid
  sequenceDiagram
    participant Client
    participant Server
    Client ->> Server: message
    Server ->> Client: message
  ```

  ```go
  message SomeRequest {}
  message SomeResponse {}

  service SomeService {
    rpc SomeMethod(SomeRequest) returns(SomeResponse)
  }
  ```

- Server Streaming RPC
  ```mermaid
  sequenceDiagram
    participant Client
    participant Server
    Client ->> Server: message
    Server ->> Client: stream message
    Server ->> Client: stream message
    Server ->> Client: stream message
  ```
  ```go
  message SomeRequest {}
  message SomeResponse {}

  service SomeService {
    rpc SomeMethod(SomeRequest) returns(stream SomeResponse)
  }
  ```

- Client Streaming RPC
  ```mermaid
  sequenceDiagram
    participant Client
    participant Server
    Client ->> Server: stream message
    Client ->> Server: stream message
    Client ->> Server: stream message
    Server ->> Client: message
  ```
  ```go
  message SomeRequest {}
  message SomeResponse {}

  service SomeService {
    rpc SomeMethod(stream SomeRequest) returns(SomeResponse)
  }
  ```

- Bidirectional Streaming RPC
  ```mermaid
  sequenceDiagram
    participant Client
    participant Server
    Client ->> Server: stream message
    Client ->> Server: stream message
    Server ->> Client: stream message
    Server ->> Client: stream message
    Server ->> Client: stream message
    Client ->> Server: stream message
  ```
  ```go
  message SomeRequest {}
  message SomeResponse {}

  service SomeService {
    rpc SomeMethod(stream SomeRequest) returns(stream SomeResponse)
  }
  ```
