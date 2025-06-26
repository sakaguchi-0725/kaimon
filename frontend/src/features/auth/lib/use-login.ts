import { useFirebaseAuth } from '@/shared/auth'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { SignInForm, signInSchema } from '../model/schemas'

export const useLogin = () => {
  const { signInWithEmailAndPassword, isLoading } = useFirebaseAuth()
  const [error, setError] = useState<string>()

  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<SignInForm>({
    resolver: zodResolver(signInSchema),
  })

  const handleLogin = handleSubmit(async (values) => {
    const { data, error } = await signInWithEmailAndPassword(
      values.email,
      values.password,
    )
    if (error) {
      setError(error)
    }
    console.log(JSON.stringify(data))
  })

  return {
    control,
    handleLogin,
    errors,
    isLoading,
    error,
  }
}
