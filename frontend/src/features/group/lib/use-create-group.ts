import { api } from '@/shared/api'
import type { CreateGroupForm } from '../model/schema'

type UseCreateGroupReturn = {
  createGroup: (data: CreateGroupForm, onSuccess?: () => void) => Promise<void>
  isLoading: boolean
  error: string | undefined
}

export const useCreateGroup = (): UseCreateGroupReturn => {
  const { mutateAsync, isPending, error } = api.useMutation('post', '/groups', {
    onError: (error) => {
      console.error('グループ作成エラー:', error)
    },
  })

  const createGroup = async (data: CreateGroupForm, onSuccess?: () => void) => {
    await mutateAsync({
      body: {
        name: data.name,
        description: data.description,
      },
    })
    onSuccess?.()
  }

  return {
    createGroup,
    isLoading: isPending,
    error:
      error?.message || (error ? 'グループの作成に失敗しました' : undefined),
  }
}
