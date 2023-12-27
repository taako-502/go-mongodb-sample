# Go MongoDB Sample

## 環境構築

```bash
# MongoDBの起動
docker run --name mongodb -d -p 27017:27017 mongo
# Echo API Serverの起動
go run app/main.go
```
