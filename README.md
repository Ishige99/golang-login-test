# 概要

- ユーザー登録時にパスワードをハッシュ化し、DBに平文保存しないようにする
- ハッシュ化したパスワードと照合するログイン機能

## 開発環境

- macOS Monterey(12.2.1)
- Goland
- go version go1.20 darwin/arm64
- Docker Compose version v2.15.1

## 使用方法

docker compose起動

```bash
% docker compose up
```

サーバー起動

```bash
% go build golang-login-test
% go run golang-login-test
```

※DB接続

```bash
% docker compose exec mysql /bin/bash
bash-4.4# mysql -u root -p
Enter password: root
```

## エンドポイント仕様

エンドポイント

- `/user`: ユーザー登録
- `/user/login`: ユーザーログイン

リクエストパラメータ

- `email`: メールアドレス
- `password`: パスワード
    - 英数字4~12文字