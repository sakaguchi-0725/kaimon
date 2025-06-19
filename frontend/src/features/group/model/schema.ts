import { validationMessages } from '@/shared/constants'
import { z } from 'zod'

export const createGroupSchema = z.object({
  name: z
    .string({ required_error: validationMessages.required })
    .min(1, { message: validationMessages.required }),
  description: z.string().optional(),
})
