import type { infer as Infer } from 'zod'
import type {
  loginSchema,
  signupSchema,
  signupConfirmSchema,
  resetPasswordSchema,
  resetPasswordConfirmSchema,
} from './schema'

export * from './schema'

export type LoginForm = Infer<typeof loginSchema>
export type SignupForm = Infer<typeof signupSchema>
export type SignupConfirmForm = Infer<typeof signupConfirmSchema>
export type ResetPasswordForm = Infer<typeof resetPasswordSchema>
export type ResetPasswordConfirmForm = Infer<typeof resetPasswordConfirmSchema>
