import type { infer as Infer } from 'zod'
import type { loginSchema } from './schema'

export * from './schema'

export type LoginForm = Infer<typeof loginSchema>
