import { api } from '@/shared/api'
import { GetAccountResponse } from '../model/types'

type UseAccountApiReturn = {
  account: GetAccountResponse | undefined
  isLoading: boolean
  error: string | undefined
  refetch: () => void
}

export const useAccountApi = (enabled: boolean = true): UseAccountApiReturn => {
  const { data, isLoading, error, refetch } = api.useQuery('get', '/account', {
    enabled,
    staleTime: 5 * 60 * 1000, // 5分間キャッシュ
  })

  return {
    account: data,
    isLoading,
    error:
      error?.message ||
      (error ? 'アカウント情報の取得に失敗しました' : undefined),
    refetch: () => refetch(),
  }
}
