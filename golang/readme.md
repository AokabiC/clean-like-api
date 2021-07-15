# Golang CleanArchitecture-like API server

[![Go tests](https://github.com/AokabiC/clean-like-api/actions/workflows/go-tests.yml/badge.svg)](https://github.com/AokabiC/clean-like-api/actions/workflows/go-tests.yml)

Create / Read / Update API

## Frameworks

- [Echo](https://echo.labstack.com/)
- [ent](https://entgo.io/)
- [gomock](https://github.com/golang/mock)
- PostgreSQL

## Run

```
docker-compose up
```

## Note
- 4つのレイヤーから構成される。
  - **domain** : アプリケーションを構成する上でのルール。DDD における値オブジェクト / エンティティ / ドメインサービスが含まれる。 
  - **usecase** : アプリケーションの機能。API を介して提供される機能(ユーザー作成 / 検索, Username の更新)がカプセル化されている。
  - **adapter**
    - 上記の層で記述される機能を REST API として提供する。HTTPRequest / Response の Validate / Format を担う。
    - データの永続化を行う(Repository の実装)。
  - **infra** : サーバーの立ち上げやDBのマイグレーションなど、アプリケーションを動かすためのコード。