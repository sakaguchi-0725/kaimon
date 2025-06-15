import { toTypedSchema } from '@vee-validate/zod'
import { useForm } from 'vee-validate'
import { loginSchema, type LoginForm } from '../model'
import { useRouter } from 'vue-router'
import { client } from '@/shared/api'

export const useLogin = () => {
  console.log('useLogin')
  console.log(import.meta.env.VITE_API_URL)
  const router = useRouter()

  const { defineField, errors, handleSubmit } = useForm<LoginForm>({
    validationSchema: toTypedSchema(loginSchema),
  })

  const onSubmit = handleSubmit(async values => {
    const { error } = await client.POST('/login', {
      body: { email: values.email, password: values.password },
    })
    if (error) return

    router.push('/groups')
  })

  return {
    defineField,
    errors,
    onSubmit,
  }
}
