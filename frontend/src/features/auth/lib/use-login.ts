import { useFirebaseAuth } from '@/shared/auth'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { SignInForm, signInSchema } from '../model/schemas'

export const useLogin = () => {
  const { signInWithEmailAndPassword, signInWithGoogle, isLoading } =
    useFirebaseAuth()
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

  const handleGoogleLogin = async () => {
    const { data, error } = await signInWithGoogle()
    if (error) {
      setError(error)
    }
    console.log(JSON.stringify(data))
  }

  return {
    control,
    handleLogin,
    handleGoogleLogin,
    errors,
    isLoading,
    error,
  }
}
