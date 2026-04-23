---
paths:
  - "docs/design/**/*.md"
  - "backend/**/*.go"
  - "backend/**/*.sql"
---

# データベース設計ガイドライン

## DDD → RDB マッピングルール

- ALWAYS Entity（集約ルート）はテーブルにする
- ALWAYS ValueObject（1:1）は同一テーブルのカラム群にする。プレフィックス付きで埋め込む（例: `address_city`, `address_zip`）
- NEVER ValueObject（1:1）を別テーブルに分離しない
- ALWAYS ValueObject（1:N）は子テーブルにする。親テーブルへのFK必須
- ALWAYS Enum はカラム（string / CHECK制約）で表現する
- NEVER Enum をマスタテーブルにしない。値の追加はマイグレーションで行う（運用で値が変わる参照データは下記「Lookup Master」を参照）
- ALWAYS 他集約への参照は FK（IDカラムのみ）で表現する
- NEVER 他集約のデータを非正規化して持たない
- ALWAYS フロントエンドに公開するIDはUUIDにする。連番IDを外部に露出しない

## Lookup Master（運用で値が変わる参照データ）

Enum に見せたいが、運用で値の追加・改名・無効化が発生する参照データは「Lookup Master」として扱う。集約化せず、専用テーブル + enum 表現の組み合わせで実装する。

- ALWAYS 運用で値が変わる参照データは専用テーブル（マスターテーブル）で管理する
- ALWAYS API / ドメイン層では enum（文字列 `code`）として表現する。ID を外部に露出しない
- NEVER Lookup Master を集約にしない（ドメインロジックを持たないため）
- ALWAYS `code` カラムに UNIQUE INDEX を付与する
- ALWAYS 他テーブルから Lookup Master を参照するときは FK（ID）で貼る
- ALWAYS repository 層で enum code → ID のサブクエリ（または JOIN）を許容する。ドメイン層に DB の ID を漏らさない
- ALWAYS 物理削除せず論理削除（`active` フラグ等）にする。過去データの `category_id` が宙に浮くのを防ぐ
- ALWAYS Go enum 値と DB マスター値の drift を検知する仕組みを入れる（アプリ起動時チェック または CI テスト）

### Enum と Lookup Master の使い分け

| 判断基準 | Enum（string + CHECK） | Lookup Master（テーブル） |
|---------|---------------------|------------------------|
| 値の変更頻度 | 低い（リリースと連動） | 高い（運用で変わる） |
| 名前変更 | マイグレーション必須 | UPDATE で完結 |
| 値の無効化 | マイグレーション必須 | `active = false` |
| ドメインロジック | enum のメソッドで完結 | 持たない前提 |

### 例

- **カテゴリー（食品 / 日用品 / ...）** → Lookup Master。運用で追加・改名が想定される
- **買い物セッションの状態（進行中 / 終了 / キャンセル）** → Enum。ドメインのライフサイクルそのものでコードと密結合

## SQLアンチパターン チェックリスト

スキーマ設計時に以下を確認する。

- NEVER カンマ区切りで複数値を1カラムに入れない（ジェイウォーク）→ 交差テーブルを使う
- NEVER 全テーブルに機械的に `id` を付けない（IDリクワイアド）→ 自然キー・複合キーを検討する
- ALWAYS FK制約を設定する（キーレスエントリ）→ 参照整合性はDB側で担保する
- NEVER attr_name / attr_value のような汎用カラムを作らない（EAV）→ 具体的なカラム定義にする
- NEVER type + id で複数テーブルを参照しない（ポリモーフィック関連）→ 交差テーブルか共通親テーブルにする
- NEVER tag1, tag2, tag3 のような横持ちカラムを作らない（マルチカラムアトリビュート）→ 従属テーブルにする
- NEVER 金額・率に FLOAT を使わない（ラウンディングエラー）→ NUMERIC / DECIMAL を使う
- NEVER 許可値を CHECK 制約にベタ書きしない（サーティワンフレーバー）→ 変更頻度が高い場合は参照テーブルを検討する
- ALWAYS 業務上必須のカラムには NOT NULL を付与する（フィア・オブ・ジ・アンノウン）

## 関連

- skill: `generate-er` -- ER図生成時
- **go-reviewer** agent -- スキーマ変更レビュー時
