# 概要

go で作った GraphQL サーバを docker の上で動かすためのレポジトリです。

M1mac 環境の人は，docker-compose.yml の db 部分に `platform: linux/amd64` を追加してして実行してください．

## 実行方法

1. レポジトリの root で `make run` する
2. intern-pj-db, intern-pj-backend の 2 つの container が立っているのを確認する．
3. `make mysql` で mysql の中にはいる．
4. `use intern-pj-development;` で intern-pj-development データベースを使用

## GraphQL サーバ について

レポジトリの root で `make start` して,[http://localhost:3000/playground](http://localhost:3000/playground)にアクセスしてください．
