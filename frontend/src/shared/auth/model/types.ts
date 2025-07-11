import { createContext } from 'react'
import { ApiSchema } from '@/shared/api'
import { FirebaseAuthTypes } from '@react-native-firebase/auth'

export type GetAccountResponse = ApiSchema<'GetAccountResponse'>

export interface AccountContextType {
  account: GetAccountResponse | undefined
  isLoading: boolean
  error: string | undefined
  refetch: () => void
}

export const AccountContext = createContext<AccountContextType | undefined>(
  undefined,
)

// Firebase認証関連の型定義
export type SignUpResponse = {
  data: FirebaseAuthTypes.User | undefined
  error: string | undefined
}

export type SignInWithEmailAndPasswordResponse = {
  data: FirebaseAuthTypes.User | undefined
  error: string | undefined
}

export type GoogleSignInResponse = {
  data: FirebaseAuthTypes.User | undefined
  error: string | undefined
}
