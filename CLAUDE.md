# Claude Code 開発ガイド

## プロジェクト概要
詳細は [README.md](./README.md) を参照してください。

## 開発ワークフロー

### 1. 実装開始
- `feature/[機能名]` ブランチを作成して作業
- `.gitmessage` テンプレートに従ってコミット
- コミット前に必ずテスト実行

### 2. 自動PR作成
実装完了後、以下のコマンドでPRを作成：
```bash
gh pr create --title "[機能名]を実装" --body-file .github/pull_request_template.md --reviewer [レビュアー名]
```

### 3. レビュー対応
PR作成後の対応手順：
- CIテストの結果を確認
- レビューコメントを受領
- 同一ブランチで修正を実施
- 修正後は追加コミットで対応

## Claude Code 固有の指示

### コミット規約
- `.gitmessage` テンプレートを必ず使用
- Conventional Commits形式を厳守
- 1コミット1機能の原則

### コーディング規約
- ESLint + Prettier設定に従う
- 全ての新機能にテストを追加
- TypeScript厳格モードを使用
- **any型の使用を絶対に禁止**（適切な型定義を必ず行う）

### 自動化コマンド
```bash
# PR作成
gh pr create --title "$(git log -1 --pretty=%s)" --body "詳細は commit message を参照"

# PR確認
gh pr view --web

# レビュー対応後のマージ
gh pr merge --squash --delete-branch
```

### 禁止事項
- main ブランチへの直接コミット
- テストなしでのPR作成
- Breaking Change の事前相談なし実施

### レビューコメント対応
レビューコメントを受けた場合：
1. 同一ブランチで修正実施
2. 修正内容を明確にコミットメッセージに記載
3. `git push` で自動的にPRに反映
4. レビュアーに修正完了を通知