import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { Alert } from 'react-native'
import { client } from '@/shared/api/client'
import { useImageUpload } from '@/shared/lib'
import { AccountInfoForm, accountInfoSchema } from '../model/schemas'
import { SignUpRequest } from '../model/types'

export const useAccountInfo = () => {
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [error, setError] = useState<string>()
  const [profileImage, setProfileImage] = useState<string>()

  const { selectFromLibrary, takePhoto, uploadImage, isUploading } =
    useImageUpload()

  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<AccountInfoForm>({
    resolver: zodResolver(accountInfoSchema),
    defaultValues: {
      accountName: '',
    },
  })

  const handleSelectImage = () => {
    Alert.alert('プロフィール画像を選択', '画像を選択してください', [
      { text: 'キャンセル', style: 'cancel' },
      {
        text: 'カメラで撮影',
        onPress: async () => {
          const imageUri = await takePhoto()
          if (imageUri) {
            setProfileImage(imageUri)
          }
        },
      },
      {
        text: 'ライブラリから選択',
        onPress: async () => {
          const imageUri = await selectFromLibrary()
          if (imageUri) {
            setProfileImage(imageUri)
          }
        },
      },
    ])
  }

  const signUp = async (data: SignUpRequest): Promise<boolean> => {
    console.log('🚀 Calling signup API with data:', data)

    const { error } = await client.POST('/signup', {
      body: data,
    })

    console.log('🔄 Signup API response:', { error })

    if (error) {
      const errorMessage = error.message || 'サインアップに失敗しました'
      console.log('❌ Signup failed with error:', errorMessage)
      setError(errorMessage)
      return false
    }

    console.log('✅ Signup successful')
    return true
  }

  const onSubmit = (
    onSuccess?: () => void,
    onError?: (errorMessage: string) => void,
  ) => {
    const submitHandler = async (data: AccountInfoForm): Promise<void> => {
      setIsSubmitting(true)
      setError(undefined)

      let uploadedImageUrl: string | null = null

      if (profileImage) {
        uploadedImageUrl = await uploadImage(profileImage)
        if (!uploadedImageUrl) {
          Alert.alert('エラー', '画像のアップロードに失敗しました')
          setIsSubmitting(false)
          return
        }
      }

      const success = await signUp({
        name: data.accountName,
        profileImageUrl: uploadedImageUrl || undefined,
      })

      if (success) {
        // Firebase Authが自動でトークンを管理するので、追加の処理は不要
        console.log('✅ Account registration completed')
        onSuccess?.()
      } else if (error) {
        onError?.(error)
      }

      setIsSubmitting(false)
    }

    return handleSubmit(submitHandler)
  }

  const isLoading = isUploading || isSubmitting

  return {
    control,
    errors,
    profileImage,
    handleSelectImage,
    onSubmit,
    isLoading,
    isUploading,
    isSubmitting,
    error,
  }
}
