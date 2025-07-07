import { z } from 'zod'

export const createGroupSchema = z.object({
  name: z
    .string({ required_error: '必須項目です' })
    .min(1, { message: '必須項目です' }),
  description: z.string().optional(),
})

export type CreateGroupForm = z.infer<typeof createGroupSchema>
