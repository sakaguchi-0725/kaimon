# グループ招待API設計

## 概要

グループへの招待機能を実装する。今回は最小限の実装として、招待コードによる招待機能を実装する。

## 実装方針

### 第1段階（今回実装）
- 6〜8桁のランダムな英数字による招待コード
- 招待コードはRedisに保持
- PostgreSQLに招待履歴テーブルは作成しない

### 第2段階（今後実装予定）
- URLによる招待機能の実装

## API仕様

### 招待コード生成API
- **Endpoint**: `POST /api/groups/{group_id}/invitations`
- **Request Body**: なし
- **Response**: 
  ```json
  {
    "invitation_code": "ABC123XY",
    "expires_at": "2024-01-01T00:00:00Z"
  }
  ```


## データ構造

### Redis
```
key: invitation:{invitation_code}
value: {
  "group_id": "group-uuid",
  "created_by": "user-uuid",
  "expires_at": "2024-01-01T00:00:00Z"
}
TTL: expires_atまでの時間
```

## 制約事項

- 招待コードの有効期限は最大7日間
- 1つのグループにつき、同時に有効な招待コードは1つまで
- 招待コードは使い回し可能（複数人が同じコードで参加可能）