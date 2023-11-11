# 概要

- protobuf について学習する

## Protocol Buffers とは

- 構造を定義するためのスキーマ言語

## なぜスキーマ言語が重要か

- マイクロサービスなどのように、複数のサービスやコンポーネントがやり取りをするためにはインタフェースについて合意しておくことで齟齬を減らすことができるから
  - Beyond the Twelve Factor Apps 等でも一番最初に出てくるのが API

## JSON との違い

- Protocol Buffers はバイナリになるため人間には分かりにくいが、型安全で軽量なのでコンピュータには優しい

## 環境構築

```bash
go mod init protobuf-lesson
```
