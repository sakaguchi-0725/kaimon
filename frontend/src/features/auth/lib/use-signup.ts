import { useForm } from 'vee-validate'
import { signupSchema, type SignupForm } from '../model'
import { toTypedSchema } from '@vee-validate/zod'
import { client } from '@/shared/api'
import { useRouter } from 'vue-router'
import { setSession } from './functions'

export const useSignup = () => {
  const router = useRouter()

  const { defineField, errors, handleSubmit } = useForm<SignupForm>({
    validationSchema: toTypedSchema(signupSchema),
  })

  const onSubmit = handleSubmit(async values => {
    const { error } = await client.POST('/signup', {
      body: {
        name: values.name,
        email: values.email,
        password: values.password,
      },
    })
    if (error) return

    setSession(values.email)
    router.push('/signup/confirm')
  })

  return {
    defineField,
    errors,
    onSubmit,
  }
}
