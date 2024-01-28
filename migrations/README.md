# Migrations

## データベースの最新化

```bash
$ make latest
```

## バージョンアップ

```bash
$ make up
```

## バージョンダウン

```bash
$ make down
```

## マイグレーションファイルの作成方法

migrations ディレクトリ配下に、`{version}_{name}.sql` というファイル名で作成する。
