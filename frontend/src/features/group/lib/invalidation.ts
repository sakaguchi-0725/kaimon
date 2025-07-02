import { queryClient } from '@/shared/api/react-query'
import { queryKeys } from './query-keys'

export const invalidateGroups = {
  all: () => queryClient.invalidateQueries({ queryKey: queryKeys.groups.all }),
  lists: () =>
    queryClient.invalidateQueries({ queryKey: queryKeys.groups.lists() }),
}
