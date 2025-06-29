import { z } from 'zod'

export const signUpSchema = z.object({
  email: z
    .string()
    .min(1, { message: '必須項目です' })
    .email({ message: 'メールアドレスの形式が不正です' }),
  password: z.string().min(1, { message: '必須項目です' }),
})

export const signInSchema = z.object({
  email: z
    .string()
    .min(1, { message: '必須項目です' })
    .email({ message: 'メールアドレスの形式が不正です' }),
  password: z.string().min(1, { message: '必須項目です' }),
})

export const accountInfoSchema = z.object({
  accountName: z
    .string()
    .min(1, { message: 'アカウント名を入力してください' })
    .min(2, { message: 'アカウント名は2文字以上で入力してください' })
    .max(20, { message: 'アカウント名は20文字以内で入力してください' }),
})

export type SignUpForm = z.infer<typeof signUpSchema>
export type SignInForm = z.infer<typeof signInSchema>
export type AccountInfoForm = z.infer<typeof accountInfoSchema>
