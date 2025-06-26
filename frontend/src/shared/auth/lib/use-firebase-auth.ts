import { useState } from 'react'
import auth from '@react-native-firebase/auth'
import { FirebaseError } from 'firebase/app'
import { SignInWithEmailAndPasswordResponse, SignUpResponse } from '../model'

export const useFirebaseAuth = () => {
  const [isLoading, setIsLoading] = useState(false)

  const signUp = async (
    email: string,
    password: string,
  ): Promise<SignUpResponse> => {
    try {
      setIsLoading(true)
      const { user } = await auth().createUserWithEmailAndPassword(
        email,
        password,
      )
      return { data: user, error: undefined }
    } catch (error) {
      if (error instanceof FirebaseError) {
        switch (error.code) {
          case 'auth/email-already-in-use':
            return {
              data: undefined,
              error: 'このメールアドレスは使用できません。',
            }
          case 'auth/invalid-email':
            return { data: undefined, error: '無効なメールアドレスです。' }
        }
      }
      return { data: undefined, error: '予期せぬエラーが発生しました。' }
    } finally {
      setIsLoading(false)
    }
  }

  const signInWithEmailAndPassword = async (
    email: string,
    password: string,
  ): Promise<SignInWithEmailAndPasswordResponse> => {
    try {
      setIsLoading(true)
      const { user } = await auth().signInWithEmailAndPassword(email, password)
      return { data: user, error: undefined }
    } catch (error) {
      if (error instanceof FirebaseError) {
        switch (error.code) {
          case 'auth/invalid-email':
            return { data: undefined, error: '無効なメールアドレスです。' }
          case 'auth/user-not-found':
            return { data: undefined, error: 'ユーザーが見つかりません。' }
        }
      }
      return { data: undefined, error: '予期せぬエラーが発生しました。' }
    } finally {
      setIsLoading(false)
    }
  }

  return {
    isLoading,
    signUp,
    signInWithEmailAndPassword,
  }
}
