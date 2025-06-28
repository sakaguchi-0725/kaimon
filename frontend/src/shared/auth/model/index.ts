import { FirebaseAuthTypes } from '@react-native-firebase/auth'

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
