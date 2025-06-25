export const AUTH_ERROR_CODES = {
  EMAIL_ALREADY_IN_USE: 'auth/email-already-in-use',
  WEAK_PASSWORD: 'auth/weak-password',
  INVALID_EMAIL: 'auth/invalid-email',
  USER_NOT_FOUND: 'auth/user-not-found',
  WRONG_PASSWORD: 'auth/wrong-password',
} as const

export const AUTH_ERROR_MESSAGES = {
  DEFAULT: '認証エラーが発生しました。しばらく時間をおいて再度お試しください。',
  WEAK_PASSWORD: 'パスワードは6文字以上で設定してください。',
  INVALID_EMAIL: 'メールアドレスの形式が正しくありません。',
  VALIDATION_ERROR: 'すべての項目を正しく入力してください。',
} as const
