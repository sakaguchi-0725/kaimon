# 画面遷移図

このドキュメントはアプリケーションの画面遷移を示しています。

## 画面一覧

- サービス概要（Welcomeページ）
- ログイン（Google認証含む）
- 会員登録（Google認証含む）
- グループ一覧
- グループ詳細
- 買い物リスト一覧
- リアルタイム買い物
- 設定
- パスワード忘れ
- 確認コード入力（パスワードリセット用）
- パスワードリセット

## 遷移図

```mermaid
flowchart TD
    Welcome["サービス概要<br/>Welcomeページ"]
    Login["ログイン<br/>Google認証含む"]
    Register["会員登録<br/>Google認証含む"]
    GroupList["グループ一覧"]
    GroupDetail["グループ詳細"]
    ShoppingLists["買い物リスト一覧"]
    ShoppingMode["リアルタイム買い物"]
    Settings["設定"]
    ForgotPw["パスワード忘れ"]
    CodePw["確認コード入力<br/>パスワードリセット用"]
    ResetPw["パスワードリセット"]
    
    Welcome --> Login
    Welcome --> Register
    Login --> GroupList
    Login --> ForgotPw
    Register --> GroupList
    
    %% グループ関連の遷移
    GroupList --> GroupDetail
    GroupList --> Settings
    GroupList --> ShoppingLists
    
    %% 買い物リスト関連の遷移
    ShoppingLists --> ShoppingMode
    
    %% パスワードリセット関連の遷移
    ForgotPw --> CodePw
    CodePw --> ResetPw
    ResetPw --> Login
``` 