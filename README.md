# Go MongoDB Sample

## 環境構築

```bash
# MongoDBの起動
docker-compose up -d
# コレクションに対してインデックスを作成
go run app/migration/create_index.go
```

### Replica Set の設定

mongo shell を開く

```bash
docker exec -it go-mongodb-sample-mongodb-1 mongosh
```

以下を実行する。

```js
rs.initiate({
  _id: "rs0",
  members: [{ _id: 0, host: "localhost:27017" }],
})
```

※
MongoDB Atlas に接続すると簡単にレプリカセットの動作確認ができるのでそれでもよい。

### API Server の起動

```bash
# Echo API Serverの起動
go run app/main.go
```
