# Kaimon フロントエンド

## 概要

Kaimonは買い物リスト管理アプリケーションのフロントエンドです。React Native（Expo）を使用したクロスプラットフォームモバイルアプリケーションとして実装されています。

## 技術スタック

### 言語・フレームワーク
- **言語**: TypeScript
- **フレームワーク**: React Native (Expo)
- **状態管理**: React Hooks

### UI/UXライブラリ
- **ナビゲーション**: React Navigation v7
- **アイコン**: react-native-feather
- **バリデーション**: react-hook-form / zod

### 認証
- Firebase Authentication
- Google認証

### その他のライブラリ
- expo-secure-store: セキュアなデータ保存
- openapi-typescript: OpenAPI定義からTypeScript型を生成

## アーキテクチャ

フロントエンドは機能ごとに分割されたFeature-Sliced Design（FSD）アーキテクチャを採用しています。

### レイヤー構成

**アプリケーション層** (`src/app/`)
- アプリケーションの初期化
- ルーティング設定
- グローバルプロバイダー

**画面層** (`src/screens/`)
- 画面コンポーネント
- 画面固有のナビゲーション

**機能層** (`src/features/`)
- 特定の機能に関するロジックとUI
- 機能ごとにディレクトリを分割

**共通層** (`src/shared/`)
- 共通UIコンポーネント
- ユーティリティ
- API通信
- 認証

### ディレクトリ構造

スライスは以下のような構造になっています：

- `model/` - 型定義、定数
- `lib/` - ロジック（hooks）
- `ui/` - UIコンポーネント

## プロジェクト構造

```
frontend/
├── android/              # Androidネイティブコード
├── ios/                  # iOSネイティブコード
├── src/
│   ├── app/              # アプリケーション層
│   │   ├── app.tsx       # アプリケーションのルート
│   │   └── navigator/    # ルートナビゲーション
│   ├── features/         # 機能モジュール
│   │   ├── auth/         # 認証機能
│   │   ├── group/        # グループ管理機能
│   │   ├── member/       # メンバー管理機能
│   │   └── shopping/     # 買い物リスト機能
│   ├── screens/          # 画面コンポーネント
│   │   ├── auth/         # 認証画面
│   │   ├── group/        # グループ画面
│   │   ├── settings/     # 設定画面
│   │   └── shopping/     # 買い物リスト画面
│   └── shared/           # 共通コンポーネント・ユーティリティ
│       ├── api/          # API通信
│       ├── auth/         # 認証ロジック
│       ├── constants/    # 定数
│       └── ui/           # 共通UIコンポーネント
├── assets/               # 画像・フォントなどの静的アセット
├── package.json          # 依存関係とスクリプト
└── tsconfig.json         # TypeScript設定
```
