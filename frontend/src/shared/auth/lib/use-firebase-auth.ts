import { useState } from 'react'
import {
  getAuth,
  createUserWithEmailAndPassword,
  signInWithEmailAndPassword as firebaseSignInWithEmailAndPassword,
  signInWithCredential,
  GoogleAuthProvider,
} from '@react-native-firebase/auth'
import { GoogleSignin } from '@react-native-google-signin/google-signin'
import { FirebaseError } from 'firebase/app'
import {
  SignInWithEmailAndPasswordResponse,
  SignUpResponse,
  GoogleSignInResponse,
} from '../model'

GoogleSignin.configure({
  webClientId: process.env.EXPO_PUBLIC_GOOGLE_WEB_CLIENT_ID,
})

export const useFirebaseAuth = () => {
  const [isLoading, setIsLoading] = useState(false)
  const auth = getAuth()

  const signUp = async (
    email: string,
    password: string,
  ): Promise<SignUpResponse> => {
    try {
      setIsLoading(true)
      const { user } = await createUserWithEmailAndPassword(
        auth,
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
          case 'auth/email-already-exists':
            return {
              data: undefined,
              error: 'このメールアドレスは使用できません。',
            }
        }
      }
      console.error(error)
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
      const { user } = await firebaseSignInWithEmailAndPassword(
        auth,
        email,
        password,
      )
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
      console.error(error)
      return { data: undefined, error: '予期せぬエラーが発生しました。' }
    } finally {
      setIsLoading(false)
    }
  }

  const signInWithGoogle = async (): Promise<GoogleSignInResponse> => {
    try {
      setIsLoading(true)
      await GoogleSignin.hasPlayServices()
      const { data } = await GoogleSignin.signIn()

      if (!data?.idToken) {
        return { data: undefined, error: 'Google認証に失敗しました。' }
      }

      const googleCredential = GoogleAuthProvider.credential(data.idToken)
      const { user } = await signInWithCredential(auth, googleCredential)

      return { data: user, error: undefined }
    } catch (error) {
      if (error instanceof FirebaseError) {
        switch (error.code) {
          case 'auth/account-exists-with-different-credential':
            return {
              data: undefined,
              error: '別の認証方法で登録されたアカウントです。',
            }
          case 'auth/invalid-credential':
            return { data: undefined, error: '認証情報が無効です。' }
        }
      }
      console.log(error)
      return { data: undefined, error: 'Google認証に失敗しました。' }
    } finally {
      setIsLoading(false)
    }
  }

  return {
    isLoading,
    signUp,
    signInWithEmailAndPassword,
    signInWithGoogle,
  }
}
