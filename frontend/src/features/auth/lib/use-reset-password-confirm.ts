import { onMounted, ref } from 'vue'
import { useForm } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/zod'
import type { ResetPasswordConfirmForm } from '../model'
import { resetPasswordConfirmSchema } from '../model'
import { getEmail, removeEmail } from './functions'
import { showError, showSuccess } from '@/shared/ui/toast'
import { useRouter } from 'vue-router'
import { client } from '@/shared/api'

export const useResetPasswordConfirm = () => {
  const router = useRouter()
  const email = ref('')

  const { defineField, errors, handleSubmit } = useForm<ResetPasswordConfirmForm>({
    validationSchema: toTypedSchema(resetPasswordConfirmSchema),
  })

  onMounted(() => {
    const value = getEmail()
    if (!value) {
      showError(
        'パスワードリセットエラー',
        'パスワードリセットの手順が正しくありません。最初からやり直してください。'
      )
      router.push({ name: 'reset-password' })
      return
    }

    email.value = value
  })

  const onSubmit = handleSubmit(async values => {
    const { error } = await client.POST('/reset-password/confirm', {
      body: {
        email: email.value,
        confirmationCode: values.code,
      },
    })
    if (error) return

    removeEmail()
    showSuccess(
      'パスワードリセット完了',
      'パスワードリセットが完了しました。ログインしてください。'
    )
    router.push({ name: 'login' })
  })

  return {
    email,
    defineField,
    errors,
    onSubmit,
  }
}
