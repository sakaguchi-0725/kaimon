import { useFirebaseAuth } from '@/shared/auth'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { SignUpForm, signUpSchema } from '../model/schemas'

export const useSignUp = () => {
  const { signUp, isLoading } = useFirebaseAuth()
  const [error, setError] = useState<string>()

  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<SignUpForm>({
    resolver: zodResolver(signUpSchema),
  })

  const handleSignUp = handleSubmit(async (values) => {
    const { data, error } = await signUp(values.email, values.password)
    if (error) {
      setError(error)
    }
    console.log(JSON.stringify(data))
  })

  return {
    control,
    handleSignUp,
    errors,
    isLoading,
    error,
  }
}
