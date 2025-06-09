# 画面遷移図

このドキュメントはアプリケーションの画面遷移を示しています。

## 画面一覧

- サービス概要（Welcomeページ）
- ログイン
- 会員登録
- 確認コード入力（会員登録用）
- グループ選択（既存グループ参加/新規グループ作成）
- グループ一覧
- グループ詳細
- 買い物リスト一覧
- 買い物モード（リアルタイム更新）
- アイテム追加/編集
- グループ作成/編集
- グループメンバー管理
- パスワード忘れ
- 確認コード入力（パスワードリセット用）
- パスワードリセット

## 遷移図

```mermaid
flowchart TD
    Welcome["サービス概要<br/>Welcomeページ"]
    Login["ログイン"]
    Register["会員登録"]
    CodeReg["確認コード入力<br/>会員登録用"]
    GroupSelect["グループ選択<br/>既存グループ参加/新規グループ作成"]
    GroupList["グループ一覧"]
    GroupDetail["グループ詳細"]
    ShoppingLists["買い物リスト一覧"]
    ShoppingMode["買い物モード<br/>リアルタイム更新"]
    ItemEdit["アイテム追加/編集"]
    GroupEdit["グループ作成/編集"]
    MemberManage["グループメンバー管理"]
    ForgotPw["パスワード忘れ"]
    CodePw["確認コード入力<br/>パスワードリセット用"]
    ResetPw["パスワードリセット"]
    
    Welcome --> Login
    Welcome --> Register
    Login --> GroupList
    Login --> ForgotPw
    Register --> CodeReg
    CodeReg --> GroupSelect
    GroupSelect --> GroupList
    
    %% グループ関連の遷移
    GroupList --> GroupDetail
    GroupList --> GroupEdit
    GroupDetail --> ShoppingLists
    GroupDetail --> MemberManage
    GroupDetail --> GroupEdit
    
    %% 買い物リスト関連の遷移
    ShoppingLists --> ItemEdit
    ShoppingLists --> ShoppingMode
    ShoppingMode --> ItemEdit
    
    %% パスワードリセット関連の遷移
    ForgotPw --> CodePw
    CodePw --> ResetPw
    ResetPw --> Login
``` 