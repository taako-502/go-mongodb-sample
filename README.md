# Go MongoDB Sample

## 環境構築

```bash
# MongoDBの起動
docker-compose up -d
# コレクションに対してインデックスを作成
go run app/migration/create_index.go
# Echo API Serverの起動
go run app/main.go
```
