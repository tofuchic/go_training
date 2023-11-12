# 概要

- protobuf について学習する

## Protocol Buffers とは

- 構造を定義するためのスキーマ言語

## なぜスキーマ言語が重要か

- マイクロサービスなどのように、複数のサービスやコンポーネントがやり取りをするためにはインタフェースについて合意しておくことで齟齬を減らすことができるから
  - Beyond the Twelve Factor Apps 等でも一番最初に出てくるのが API

## JSON との違い

- Protocol Buffers はバイナリになるため人間には分かりにくいが、型安全で軽量なのでコンピュータには優しい

## ワーク内容

- Goプロジェクトの初期化
  ```bash
  go mod init protobuf-lesson
  ```

- スキーマであるprotoファイルの作成
  - proto/*.proto

- go言語での型定義生成
  ```bash
  protoc -I. --go_out=. proto/*.proto
  ```
    - これにより、`option go_package`で指定したディレクトリ名が生成され、その配下にgoファイルが生成

- 生成されたコードに必要なモジュールのインストール
  ```bash
  go mod tidy
  ```

- 型定義を使うスクリプトの作成
  - main.goの作成

- main.goの実行
  ```bash
  go run main.go
  ```