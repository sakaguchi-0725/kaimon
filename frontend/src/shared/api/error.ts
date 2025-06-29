import { Alert } from 'react-native'
import type { components } from './schema'

export type ApiError = components['schemas']['Error']

export const isApiError = (error: unknown): error is ApiError => {
  return (
    typeof error === 'object' &&
    error !== null &&
    'code' in error &&
    'message' in error &&
    typeof (error as any).code === 'string' &&
    typeof (error as any).message === 'string'
  )
}

export const handleUnexpectedError = (error: unknown) => {
  console.error('Unexpected API error:', error)
  Alert.alert('エラー', 'ネットワークエラーが発生しました')
}
