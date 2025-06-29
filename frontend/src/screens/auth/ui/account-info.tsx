import React from 'react'
import {
  StyleSheet,
  Text,
  View,
  TouchableOpacity,
  Image,
  Alert,
} from 'react-native'
import { Controller } from 'react-hook-form'
import { Container } from '@/shared/ui/container'
import { TextInput } from '@/shared/ui/input'
import { Button } from '@/shared/ui/button'
import { Colors } from '@/shared/constants'
import { useNavigation } from '@react-navigation/native'
import { useAccountInfo } from '@/features/auth'
import { AuthNavigationProp } from '@/screens/auth'

export const AccountInfoScreen = () => {
  const navigation = useNavigation<AuthNavigationProp>()
  const {
    control,
    errors,
    profileImage,
    handleSelectImage,
    onSubmit,
    isLoading,
  } = useAccountInfo()

  const handleAccountSubmit = onSubmit(
    () => {
      // アクセストークンが保存されることで自動的にグループ一覧画面に遷移
      Alert.alert('成功', 'アカウント登録が完了しました')
    },
    (errorMessage) => {
      Alert.alert('エラー', errorMessage)
    },
  )

  return (
    <Container style={styles.container}>
      <View style={styles.header}>
        <Text style={styles.title}>アカウント情報を入力</Text>
        <Text style={styles.subtitle}>プロフィールを設定しましょう</Text>
      </View>

      <View style={styles.form}>
        <View style={styles.profileImageContainer}>
          <TouchableOpacity
            style={styles.profileImageButton}
            onPress={handleSelectImage}
          >
            {profileImage ? (
              <Image
                source={{ uri: profileImage }}
                style={styles.profileImage}
              />
            ) : (
              <View style={styles.placeholderImage}>
                <Text style={styles.placeholderText}>写真を追加</Text>
              </View>
            )}
          </TouchableOpacity>
          <Text style={styles.imageHint}>プロフィール画像を設定（任意）</Text>
        </View>

        <Controller
          control={control}
          name="accountName"
          render={({ field: { onChange, onBlur, value } }) => (
            <TextInput
              label="アカウント名"
              placeholder="アカウント名を入力"
              required
              containerStyle={styles.inputContainer}
              value={value}
              onChangeText={onChange}
              onBlur={onBlur}
              error={errors.accountName?.message}
              maxLength={20}
            />
          )}
        />

        <Button
          text={isLoading ? '登録中...' : '完了'}
          onPress={handleAccountSubmit}
          color="primary"
          style={styles.button}
          disabled={isLoading}
        />
      </View>

      <View style={styles.footer}>
        <TouchableOpacity onPress={() => navigation.goBack()}>
          <Text style={styles.link}>戻る</Text>
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
    alignItems: 'center',
  },
  title: {
    fontSize: 28,
    fontWeight: 'bold',
    color: Colors.mainText,
    marginBottom: 8,
    textAlign: 'center',
  },
  subtitle: {
    fontSize: 16,
    color: Colors.subText,
    textAlign: 'center',
  },
  form: {
    marginBottom: 24,
  },
  profileImageContainer: {
    alignItems: 'center',
    marginBottom: 32,
  },
  profileImageButton: {
    width: 120,
    height: 120,
    borderRadius: 60,
    marginBottom: 12,
  },
  profileImage: {
    width: 120,
    height: 120,
    borderRadius: 60,
  },
  placeholderImage: {
    width: 120,
    height: 120,
    borderRadius: 60,
    borderWidth: 2,
    borderColor: Colors.border,
    borderStyle: 'dashed',
    justifyContent: 'center',
    alignItems: 'center',
  },
  placeholderText: {
    color: Colors.subText,
    fontSize: 14,
    textAlign: 'center',
  },
  imageHint: {
    color: Colors.subText,
    fontSize: 12,
    textAlign: 'center',
  },
  inputContainer: {
    marginBottom: 16,
  },
  button: {
    marginTop: 24,
  },
  footer: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
  },
  link: {
    color: Colors.primary,
    fontWeight: 'bold',
  },
})
