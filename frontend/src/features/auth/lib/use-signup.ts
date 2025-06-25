import { Alert } from 'react-native'
import { useNavigation } from '@react-navigation/native'
import { NativeStackNavigationProp } from '@react-navigation/native-stack'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { FirebaseError } from 'firebase/app'
import { signUp } from '../api'
import { signUpSchema, SignUpFormData } from './schemas'
import { AUTH_ERROR_CODES, AUTH_ERROR_MESSAGES } from '../model'

type AuthStackParamList = {
  Login: undefined
  SignUp: undefined
}

type NavigationProp = NativeStackNavigationProp<AuthStackParamList>

export const useSignup = () => {
  const navigation = useNavigation<NavigationProp>()

  const {
    control,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<SignUpFormData>({
    resolver: zodResolver(signUpSchema),
    defaultValues: {
      userName: '',
      email: '',
      password: '',
    },
  })

  const handleSignUp = handleSubmit(async (data) => {
    try {
      const user = await signUp(data.email, data.password)
      console.log('登録成功。UID:', user.uid)
      Alert.alert('成功', '会員登録が完了しました', [
        { text: 'OK', onPress: () => navigation.navigate('Login') },
      ])
    } catch (error) {
      let errorMessage = AUTH_ERROR_MESSAGES.DEFAULT

      if (error instanceof FirebaseError) {
        switch (error.code) {
          case AUTH_ERROR_CODES.WEAK_PASSWORD:
            errorMessage = AUTH_ERROR_MESSAGES.WEAK_PASSWORD
            break
          case AUTH_ERROR_CODES.INVALID_EMAIL:
            errorMessage = AUTH_ERROR_MESSAGES.INVALID_EMAIL
            break
          case AUTH_ERROR_CODES.EMAIL_ALREADY_IN_USE:
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
    handleSignUp,
  }
}
