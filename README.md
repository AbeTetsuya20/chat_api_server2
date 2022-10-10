# 仕様

友達と作る、チャットゲームの API サーバーです。
sendai.go の Gocon で発表予定。
クリーンアーキテクチャで実装しました。

# API 一覧

## 権限なし
以下の API は User トークンや Admin トークンなしで叩けます。

### /api/users
ユーザーの一覧の取得

- 実行例
```shell
curl http://localhost:1001/api/users
```

- 期待される結果
```json
{
  "Users": [
    {
      "name": "テストユーザー",
      "ID": "test_1234",
      "Token": {
        "String": "test_1234vl5",
        "Valid": true
      },
      "chatNumber": 10
    },
    {
      "name": "テストユーザー",
      "ID": "test_5678",
      "Token": {
        "String": "test_5678_1234",
        "Valid": true
      },
      "chatNumber": 10
    }
  ]
}
```

### /api/signup
新規ユーザーの登録

- 実行例
```shell
curl -H 'name:name3' -H 'address:tmp@tmp.com' http://localhost:1001/api/signup
```

- 期待される結果

```json
{
  "success": true
}
```

### /api/login/user
一般ユーザーとしてログイン

- 実行例
```shell
curl -H 'name:テストユーザー' -H 'address:test1@test.com' -H 'password:test' http://localhost:1001/api/login/user
```

- 期待される結果

```json
{
	"success": true,
	"id": "test_1234",
	"token": {
		"String": "test_1234adM",
		"Valid": false
	}
}
```

### /api/login/admin
アドミンとしてログイン

- 実行例
```shell
curl -H 'id:admin' -H 'password:ABC123abc' http://localhost:1001/api/login/admin
```

- 期待される結果

```json
{
  "success": true,
  "id": "admin1_12345",
  "token": "token_admin1_12345abcdifgABC"
}
```

## 権限あり: User トークン
以下の API はログイン時に帰ってくる User トークンが必要です。

### /api/user/profile
User プロフィールの変更
- 実行例

```shell
BODY='{
  "id":"user_xxx",
  "profile_message": "こんにちは！編集済です！"
}'

curl -i \
  -H 'token: token_name1_12345abcdifgABC' \
  -H 'Content-Type: application/json' \
  -d "$BODY" \
  http://localhost:1001/api/user/profile
```

- 期待される結果

```json
{
  "success": true
}
```

## 権限あり: Admin トークン
以下の API は Admin でログイン時に帰ってくる Admin トークンが必要です。

### /api/admin/ban
特定のユーザーを ban できます。

- 実行例

```shell
BODY='{
  "user_id": "token_name1_12345abcdifgABC",
  "reason": "不正なログイン"
}'

curl -i \
  -H 'token: token_admin1_12345abcdifgABC' \
  -d "$BODY" \
  http://localhost:1001/api/admin/ban
```

- 期待される結果

```json
{
  "success": true,
  "error": null
}
```