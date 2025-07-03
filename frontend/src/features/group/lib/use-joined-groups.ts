import { api } from '@/shared/api'
import { queryKeys } from './query-keys'
import type { JoinedGroup } from '../model'

type UseJoinedGroupsReturn = {
  groups: JoinedGroup[]
  isLoading: boolean
  error: string | undefined
  refetch: () => void
}

export const useJoinedGroups = (): UseJoinedGroupsReturn => {
  const { data, isLoading, error, refetch } = api.useQuery('get', '/groups', {
    queryKey: queryKeys.groups.lists(),
    staleTime: 5 * 60 * 1000, // 5分間キャッシュ
  })

  return {
    groups: data?.groups || [],
    isLoading,
    error:
      error?.message ||
      (error ? 'グループ一覧の取得に失敗しました' : undefined),
    refetch: () => refetch(),
  }
}
