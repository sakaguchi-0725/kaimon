import { z } from 'zod'

export const signUpSchema = z.object({
  name: z.string().min(1, { message: '必須項目です' }),
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

export type SignUpForm = z.infer<typeof signUpSchema>
export type SignInForm = z.infer<typeof signInSchema>
