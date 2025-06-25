import React, { useState } from 'react'
import { StyleSheet, Text, View, TouchableOpacity, Alert } from 'react-native'
import { Container } from '@/shared/ui/container'
import { TextInput } from '@/shared/ui/input'
import { Button } from '@/shared/ui/button'
import { Colors } from '@/shared/constants'
import { useNavigation } from '@react-navigation/native'
import { NativeStackNavigationProp } from '@react-navigation/native-stack'
import { GoogleIcon } from '@/shared/ui/icons'
import { authService } from '@/shared/auth'

type AuthStackParamList = {
  Login: undefined
  SignUp: undefined
}

type NavigationProp = NativeStackNavigationProp<AuthStackParamList>

export const SignUpScreen = () => {
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

  const handleGoogleSignUp = () => {
    // Googleサインアップ処理（実装不要）
  }

  const navigateToLogin = () => {
    navigation.navigate('Login')
  }

  return (
    <Container style={styles.container}>
      <View style={styles.header}>
        <Text style={styles.title}>会員登録</Text>
        <Text style={styles.subtitle}>アカウントを作成して始めましょう</Text>
      </View>

      <View style={styles.form}>
        <TextInput
          label="ユーザー名"
          placeholder="ユーザー名を入力"
          value={userName}
          onChangeText={setUserName}
          required
          containerStyle={styles.inputContainer}
        />

        <TextInput
          label="メールアドレス"
          placeholder="メールアドレスを入力"
          value={email}
          onChangeText={setEmail}
          keyboardType="email-address"
          autoCapitalize="none"
          required
          containerStyle={styles.inputContainer}
        />

        <TextInput
          label="パスワード"
          placeholder="パスワードを入力"
          value={password}
          onChangeText={setPassword}
          secureTextEntry
          required
          containerStyle={styles.inputContainer}
        />

        <Button
          text="会員登録"
          onPress={handleSignUp}
          color="primary"
          style={styles.button}
          disabled={isLoading}
        />

        <View style={styles.dividerContainer}>
          <View style={styles.divider} />
          <Text style={styles.dividerText}>または</Text>
          <View style={styles.divider} />
        </View>

        <Button
          text="Googleで登録"
          onPress={handleGoogleSignUp}
          variant="outline"
          color="secondary"
          style={styles.googleButton}
          textStyle={styles.googleButtonText}
          icon={<GoogleIcon size={20} />}
          iconPosition="left"
        />
      </View>

      <View style={styles.footer}>
        <Text style={styles.footerText}>すでにアカウントをお持ちですか？</Text>
        <TouchableOpacity onPress={navigateToLogin}>
          <Text style={styles.link}>ログイン</Text>
        </TouchableOpacity>
      </View>
    </Container>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    backgroundColor: Colors.white,
    padding: 24,
  },
  header: {
    marginBottom: 32,
  },
  title: {
    fontSize: 28,
    fontWeight: 'bold',
    color: Colors.mainText,
    marginBottom: 8,
  },
  subtitle: {
    fontSize: 16,
    color: Colors.subText,
  },
  form: {
    marginBottom: 24,
  },
  inputContainer: {
    marginBottom: 16,
  },
  button: {
    marginTop: 8,
  },
  dividerContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    marginVertical: 24,
  },
  divider: {
    flex: 1,
    height: 1,
    backgroundColor: Colors.border,
  },
  dividerText: {
    color: Colors.subText,
    paddingHorizontal: 16,
    fontSize: 14,
  },
  googleButton: {
    backgroundColor: Colors.white,
    borderColor: Colors.border,
  },
  googleButtonText: {
    color: Colors.mainText,
  },
  footer: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
  },
  footerText: {
    color: Colors.subText,
    marginRight: 4,
  },
  link: {
    color: Colors.primary,
    fontWeight: 'bold',
  },
})
