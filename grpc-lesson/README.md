# 概要

- gRPCについて学習する

## gRPCとは

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

## 4つの通信方式とServiceの記述例

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

## Interceptor

- 認証認可など、全てのメソッドに共通して行いたい処理を実装できる機能
  - 決まった型で実装し、サーバサイドおよびクライアントサイドでそれぞれ使用できる
    - 型
      ```go
      type UnaryServerInterceptor func(
        ctx context.Context,
        req interface{},
        info *UnaryServerInfo, // リクエストを受け付けたメソッド名などのサーバサイドの情報
        handler UnaryHandler, // リクエストを受け付けたメソッド
      ) (resp interface{}, err error)
      ```
    - サーバサイドへの埋め込み
      ```go
      s := grpc.NewServer(
        grpc.UnaryInterceptor(myInterceptor()),
      )
      ```
    - クライアントサイドへの埋め込み
      ```go
      conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithUnaryInterceptor(myInterceptor()));
      ```


## ワーク内容

### protoファイルの作成まで

- Goプロジェクトの初期化
  ```bash
  go mod init grpc-lesson
  ```

- スキーマであるprotoファイルの作成
  - proto/*.proto

- go言語でのgRPC用の型定義生成
  ```bash
  protoc -I. --go_out=. --go-grpc_out=. proto/*.proto
  ```
    - これにより、`option go_package`で指定したディレクトリ名が生成され、その配下にgoファイルが生成

- 生成されたコードに必要なモジュールのインストール
  ```bash
  go mod tidy
  ```

### サーバサイドの実装

- 型定義を使うスクリプトの作成
  - server/main.goの作成

- main.goの実行
  ```bash
  go run main.go
  ```

### クライアントサイドの実装

- 型定義を使うスクリプトの作成
  - client/main.goの作成

- main.goの実行
  ```bash
  go run main.go
  ```
