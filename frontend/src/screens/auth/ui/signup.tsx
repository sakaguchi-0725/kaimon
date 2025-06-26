import React from 'react'
import { StyleSheet, Text, View, TouchableOpacity } from 'react-native'
import { Controller } from 'react-hook-form'
import { Container } from '@/shared/ui/container'
import { TextInput } from '@/shared/ui/input'
import { Button } from '@/shared/ui/button'
import { Colors } from '@/shared/constants'
import { useNavigation } from '@react-navigation/native'
import { NativeStackNavigationProp } from '@react-navigation/native-stack'
import { GoogleIcon } from '@/shared/ui/icons'
import { useSignUp } from '@/features/auth'

type AuthStackParamList = {
  Login: undefined
  SignUp: undefined
}

type NavigationProp = NativeStackNavigationProp<AuthStackParamList>

export const SignUpScreen = () => {
  const navigation = useNavigation<NavigationProp>()
  const { control, errors, handleSignUp, isLoading, error } = useSignUp()

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
        <Controller
          control={control}
          name="name"
          render={({ field: { onChange, onBlur, value } }) => (
            <TextInput
              label="ユーザー名"
              placeholder="ユーザー名を入力"
              required
              containerStyle={styles.inputContainer}
              value={value}
              onChangeText={onChange}
              onBlur={onBlur}
              error={errors.name?.message}
            />
          )}
        />

        <Controller
          control={control}
          name="email"
          render={({ field: { onChange, onBlur, value } }) => (
            <TextInput
              label="メールアドレス"
              placeholder="メールアドレスを入力"
              keyboardType="email-address"
              autoCapitalize="none"
              required
              containerStyle={styles.inputContainer}
              value={value}
              onChangeText={onChange}
              onBlur={onBlur}
              error={errors.email?.message}
            />
          )}
        />

        <Controller
          control={control}
          name="password"
          render={({ field: { onChange, onBlur, value } }) => (
            <TextInput
              label="パスワード"
              placeholder="パスワードを入力"
              secureTextEntry
              required
              containerStyle={styles.inputContainer}
              value={value}
              onChangeText={onChange}
              onBlur={onBlur}
              error={errors.password?.message}
            />
          )}
        />

        {error && (
          <View style={styles.errorContainer}>
            <Text style={styles.errorText}>{error}</Text>
          </View>
        )}

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
  errorContainer: {
    marginBottom: 16,
    padding: 12,
    backgroundColor: Colors.error + '10',
    borderRadius: 8,
    borderWidth: 1,
    borderColor: Colors.error + '30',
  },
  errorText: {
    color: Colors.error,
    fontSize: 14,
    textAlign: 'center',
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
