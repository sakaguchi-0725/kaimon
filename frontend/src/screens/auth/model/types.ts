import { NativeStackNavigationProp } from '@react-navigation/native-stack'

export type AuthStackParamList = {
  Welcome: undefined
  Login: undefined
  SignUp: undefined
  AccountInfo: undefined
}

export type AuthNavigationProp = NativeStackNavigationProp<AuthStackParamList>
