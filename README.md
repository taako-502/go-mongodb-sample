# Go MongoDB Sample

[![codecov](https://codecov.io/gh/taako-502/go-mongodb-sample/graph/badge.svg?token=1R0OXO8EO4)](https://codecov.io/gh/taako-502/go-mongodb-sample)
[![Go Report Card](https://goreportcard.com/badge/github.com/taako-502/go-mongodb-sample)](https://goreportcard.com/report/github.com/taako-502/go-mongodb-sample)

## 環境構築

```bash
openssl rand -base64 756 > mongodb-keyfile
chmod 400 mongodb-keyfile
docker-compose build
docker-compose up -d
# コレクションに対してインデックスを作成
go run app/migration/create_index.go
```

### ローカル環境 の Replica Set の設定

mongo shell を開く。

```bash
docker exec -it go_mongodb_sample_db mongosh
```

以下を実行する。

```js
rs.initiate({
  _id: "rs0",
  members: [{ _id: 0, host: "go_mongodb_sample_db:27017" }],
})
```

```bash
rs.status()
rs.initiate()
```

testdb に接続する方法。
レプリカの設定を以下のコマンドで行う必要がある。

```bash
var config = rs.conf();
config.members[0].host = "mongo_db:27017";
rs.reconfig(config, { force: true });
```

※

MongoDB Atlas に接続すると簡単にレプリカセットの動作確認ができるのでそれでもよい。

### CI 用のデータベースの設定

mongo shell を開く。

```bash
docker exec -it go_mongodb_sample_db_ci mongosh
```

以下を実行する。

```js
rs.initiate({
  _id: "rs0",
  members: [{ _id: 0, host: "go_mongodb_sample_db_ci:27018" }],
})
```

### API Server の起動

```bash
# Echo API Serverの起動
go run app/main.go
```
