import React, { useState } from 'react'
import {
  StyleSheet,
  Text,
  View,
  TouchableOpacity,
  Image,
  Alert,
} from 'react-native'
import { Controller, useForm } from 'react-hook-form'
import { Container } from '@/shared/ui/container'
import { TextInput } from '@/shared/ui/input'
import { Button } from '@/shared/ui/button'
import { Colors } from '@/shared/constants'
import { useNavigation } from '@react-navigation/native'
import { NativeStackNavigationProp } from '@react-navigation/native-stack'
import { AuthStackParamList } from './stack-navigator'

type NavigationProp = NativeStackNavigationProp<AuthStackParamList>

interface AccountInfoForm {
  accountName: string
}

export const AccountInfoScreen = () => {
  const navigation = useNavigation<NavigationProp>()
  const [profileImage, setProfileImage] = useState<string | null>(null)
  const [isLoading, setIsLoading] = useState(false)

  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<AccountInfoForm>({
    defaultValues: {
      accountName: '',
    },
  })

  const handleSelectImage = () => {
    Alert.alert('プロフィール画像を選択', '画像を選択してください', [
      { text: 'キャンセル', style: 'cancel' },
      { text: 'カメラで撮影', onPress: () => {} },
      { text: 'ライブラリから選択', onPress: () => {} },
    ])
  }

  const handleComplete = (data: AccountInfoForm) => {
    setIsLoading(true)

    setTimeout(() => {
      setIsLoading(false)
      console.log('Account info:', { ...data, profileImage })
    }, 1000)
  }

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
          rules={{
            required: 'アカウント名を入力してください',
            minLength: {
              value: 2,
              message: 'アカウント名は2文字以上で入力してください',
            },
            maxLength: {
              value: 20,
              message: 'アカウント名は20文字以内で入力してください',
            },
          }}
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
          text="完了"
          onPress={handleSubmit(handleComplete)}
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
