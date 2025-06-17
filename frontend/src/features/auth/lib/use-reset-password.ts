import { useForm } from 'vee-validate'
import { useRouter } from 'vue-router'
import { resetPasswordSchema, type ResetPasswordForm } from '../model'
import { toTypedSchema } from '@vee-validate/zod'
import { client } from '@/shared/api'
import { showSuccess } from '@/shared/ui/toast'
import { setEmail } from './functions'

export const useResetPassword = () => {
  const router = useRouter()
  const { defineField, errors, handleSubmit } = useForm<ResetPasswordForm>({
    validationSchema: toTypedSchema(resetPasswordSchema),
  })

  const onSubmit = handleSubmit(async values => {
    const { error } = await client.POST('/reset-password', {
      body: { email: values.email },
    })

    if (error) return

    showSuccess('パスワードリセット', 'パスワードリセットのメールを送信しました。')
    setEmail(values.email)
    router.push({ name: 'reset-password-confirm' })
  })

  return {
    defineField,
    errors,
    onSubmit,
  }
}
