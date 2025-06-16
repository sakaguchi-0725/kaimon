import type { infer as Infer } from 'zod'
import type { loginSchema, signupSchema } from './schema'

export * from './schema'

export type LoginForm = Infer<typeof loginSchema>
export type SignupForm = Infer<typeof signupSchema>
