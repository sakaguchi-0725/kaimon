import { useForm } from 'vee-validate'
import { signupConfirmSchema, type SignupConfirmForm } from '../model'
import { onMounted, ref } from 'vue'
import { getEmail, removeEmail } from './functions'
import { showError, showSuccess } from '@/shared/ui/toast'
import { useRouter } from 'vue-router'
import { client } from '@/shared/api'

export const useSignupConfirm = () => {
  const router = useRouter()
  const email = ref('')

  const { defineField, errors, handleSubmit } = useForm<SignupConfirmForm>({
    validationSchema: signupConfirmSchema,
  })

  onMounted(() => {
    const value = getEmail()
    if (!value) {
      showError('登録エラー', '登録の手順が正しくありません。最初からやり直してください。')
      router.push({ name: 'signup' })
      return
    }

    email.value = value
  })

  const onSubmit = handleSubmit(async values => {
    const { error } = await client.POST('/signup/confirm', {
      body: {
        email: email.value,
        confirmationCode: values.code,
      },
    })
    if (error) return

    removeEmail()
    showSuccess('登録完了', 'アカウント登録が完了しました。ログインしてください。')
    router.push({ name: 'login' })
  })

  return {
    email,
    defineField,
    errors,
    onSubmit,
  }
}
