import { useState } from 'react'
import { Alert } from 'react-native'
import { useNavigation } from '@react-navigation/native'
import { NativeStackNavigationProp } from '@react-navigation/native-stack'
import { authService } from '@/shared/lib/firebase'

type AuthStackParamList = {
  Login: undefined
  SignUp: undefined
}

type NavigationProp = NativeStackNavigationProp<AuthStackParamList>

export const useSignup = () => {
  const navigation = useNavigation<NavigationProp>()
  const [userName, setUserName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [isLoading, setIsLoading] = useState(false)

  const handleSignUp = async () => {
    if (!userName.trim() || !email.trim() || !password.trim()) {
      Alert.alert('エラー', 'すべての項目を入力してください')
      return
    }

    setIsLoading(true)
    try {
      const user = await authService.signUp(email, password)
      console.log('登録成功。UID:', user.uid)
      Alert.alert('成功', '会員登録が完了しました', [
        { text: 'OK', onPress: () => navigation.navigate('Login') },
      ])
    } catch (error: any) {
      let errorMessage = '会員登録に失敗しました'

      if (error.code === 'auth/email-already-in-use') {
        errorMessage = 'このメールアドレスは既に使用されています'
      } else if (error.code === 'auth/weak-password') {
        errorMessage = 'パスワードが弱すぎます'
      } else if (error.code === 'auth/invalid-email') {
        errorMessage = 'メールアドレスの形式が正しくありません'
      }

      Alert.alert('エラー', errorMessage)
    } finally {
      setIsLoading(false)
    }
  }

  return {
    userName,
    setUserName,
    email,
    setEmail,
    password,
    setPassword,
    isLoading,
    handleSignUp,
  }
}
