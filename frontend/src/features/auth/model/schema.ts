import { validationMessages } from '@/shared/constants'
import { z } from 'zod'

export const loginSchema = z.object({
  email: z
    .string({ required_error: validationMessages.required })
    .email({ message: validationMessages.email })
    .min(1, { message: validationMessages.required }),
  password: z
    .string({ required_error: validationMessages.required })
    .min(1, { message: validationMessages.required }),
})

export const signupSchema = z.object({
  name: z
    .string({ required_error: validationMessages.required })
    .min(1, { message: validationMessages.required }),
  email: z
    .string({ required_error: validationMessages.required })
    .email({ message: validationMessages.email })
    .min(1, { message: validationMessages.required }),
  password: z
    .string({ required_error: validationMessages.required })
    .min(1, { message: validationMessages.required }),
})

export const signupConfirmSchema = z.object({
  code: z
    .string({ required_error: validationMessages.required })
    .min(1, { message: validationMessages.required }),
})

export const resetPasswordSchema = z.object({
  email: z
    .string({ required_error: validationMessages.required })
    .email({ message: validationMessages.email })
    .min(1, { message: validationMessages.required }),
})

export const resetPasswordConfirmSchema = z.object({
  code: z
    .string({ required_error: validationMessages.required })
    .min(1, { message: validationMessages.required }),
})
