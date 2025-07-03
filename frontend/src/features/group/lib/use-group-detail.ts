import { api } from '@/shared/api'
import { GroupInfo } from '../model'

type UseGroupDetailResult = {
  data: GroupInfo | undefined
  isLoading: boolean
  error: string | undefined
  refetch: () => Promise<void>
}

export const useGroupDetail = (groupId: string): UseGroupDetailResult => {
  const query = api.useQuery('get', '/groups/{id}', {
    params: {
      path: { id: groupId },
    },
  })

  return {
    data: query.data,
    isLoading: query.isLoading,
    error: query.error?.message,
    refetch: async () => {
      await query.refetch()
    },
  }
}
