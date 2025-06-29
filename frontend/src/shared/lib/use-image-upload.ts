import { useState } from 'react'
import * as ImagePicker from 'expo-image-picker'
import {
  getStorage,
  ref,
  putFile,
  getDownloadURL,
} from '@react-native-firebase/storage'
import { Alert } from 'react-native'
import { getAuth } from '@react-native-firebase/auth'

export const useImageUpload = () => {
  const [isUploading, setIsUploading] = useState(false)

  const requestPermissions = async () => {
    const { status } = await ImagePicker.requestMediaLibraryPermissionsAsync()
    if (status !== 'granted') {
      Alert.alert('権限エラー', '写真ライブラリへのアクセス権限が必要です')
      return false
    }
    return true
  }

  const requestCameraPermissions = async () => {
    const { status } = await ImagePicker.requestCameraPermissionsAsync()
    if (status !== 'granted') {
      Alert.alert('権限エラー', 'カメラへのアクセス権限が必要です')
      return false
    }
    return true
  }

  const selectFromLibrary = async (): Promise<string | null> => {
    const hasPermission = await requestPermissions()
    if (!hasPermission) return null

    const result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ['images'],
      allowsEditing: true,
      aspect: [1, 1],
      quality: 0.8,
    })

    if (!result.canceled && result.assets[0]) {
      return result.assets[0].uri
    }
    return null
  }

  const takePhoto = async (): Promise<string | null> => {
    const hasPermission = await requestCameraPermissions()
    if (!hasPermission) return null

    const result = await ImagePicker.launchCameraAsync({
      allowsEditing: true,
      aspect: [1, 1],
      quality: 0.8,
    })

    if (!result.canceled && result.assets[0]) {
      return result.assets[0].uri
    }
    return null
  }

  const uploadImage = async (imageUri: string): Promise<string | null> => {
    if (!imageUri) return null

    const auth = getAuth()
    const user = auth.currentUser
    if (!user) {
      throw new Error('ユーザーが認証されていません')
    }

    setIsUploading(true)
    try {
      const filename = `profile_images/${user.uid}_${Date.now()}.jpg`
      const storage = getStorage()
      const reference = ref(storage, filename)

      await putFile(reference, imageUri)
      const downloadURL = await getDownloadURL(reference)

      return downloadURL
    } catch (error) {
      console.error('画像アップロードエラー:', error)
      Alert.alert('エラー', '画像のアップロードに失敗しました')
      return null
    } finally {
      setIsUploading(false)
    }
  }

  return {
    isUploading,
    selectFromLibrary,
    takePhoto,
    uploadImage,
  }
}
