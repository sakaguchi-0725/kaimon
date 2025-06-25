import { Alert } from 'react-native'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { FirebaseError } from 'firebase/app'
import { signIn } from '../api'
import { loginSchema, LoginFormData } from '../model/schemas'
import { AUTH_ERROR_CODES, AUTH_ERROR_MESSAGES } from '../model'
import { authStorage } from '@/shared/lib/auth-storage'

export const useLogin = () => {
  const {
    control,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<LoginFormData>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: '',
      password: '',
    },
  })

  const handleLogin = handleSubmit(async (data) => {
    try {
      const user = await signIn(data.email, data.password)

      // AccessTokenとRefreshTokenを取得して保存
      const accessToken = await user.getIdToken()
      const refreshToken = user.refreshToken

      await authStorage.setTokens(accessToken, refreshToken)

      console.log('ログイン成功。UID:', user.uid)
      console.log('AccessToken保存完了')
      console.log('RefreshToken保存完了')

      // TODO: ログイン成功後の画面遷移を実装
      Alert.alert('成功', 'ログインが完了しました')
    } catch (error) {
      let errorMessage = AUTH_ERROR_MESSAGES.DEFAULT

      if (error instanceof FirebaseError) {
        switch (error.code) {
          case AUTH_ERROR_CODES.INVALID_EMAIL:
            errorMessage = AUTH_ERROR_MESSAGES.INVALID_EMAIL
            break
          case AUTH_ERROR_CODES.USER_NOT_FOUND:
          case AUTH_ERROR_CODES.WRONG_PASSWORD:
          default:
            errorMessage = AUTH_ERROR_MESSAGES.DEFAULT
            break
        }
      }

      Alert.alert('エラー', errorMessage)
    }
  })

  return {
    control,
    errors,
    isLoading: isSubmitting,
    handleLogin,
  }
}
