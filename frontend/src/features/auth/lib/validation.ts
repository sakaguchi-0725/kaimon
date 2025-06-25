import { z } from 'zod'

export const signUpSchema = z.object({
  userName: z
    .string()
    .min(1, 'ユーザー名を入力してください')
    .max(50, 'ユーザー名は50文字以内で入力してください'),
  email: z
    .string()
    .min(1, 'メールアドレスを入力してください')
    .email('メールアドレスの形式が正しくありません'),
  password: z
    .string()
    .min(6, 'パスワードは6文字以上で入力してください')
    .max(100, 'パスワードは100文字以内で入力してください'),
})

export type SignUpFormData = z.infer<typeof signUpSchema>
