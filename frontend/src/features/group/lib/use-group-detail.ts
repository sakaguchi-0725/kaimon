import { useState, useEffect } from 'react'
import { client } from '@/shared/api'
import { GroupInfo } from '../model'

interface UseGroupDetailResult {
  data: GroupInfo | null
  isLoading: boolean
  error: string | null
  refetch: () => Promise<void>
}

export const useGroupDetail = (groupId: string): UseGroupDetailResult => {
  const [data, setData] = useState<GroupInfo | null>(null)
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const fetchGroupDetail = async () => {
    setIsLoading(true)
    setError(null)

    try {
      const response = await client.GET('/groups/{id}', {
        params: {
          path: { id: groupId },
        },
      })

      if (response.error) {
        setError('グループ詳細の取得に失敗しました')
      } else {
        setData(response.data)
      }
    } catch (err) {
      setError('グループ詳細の取得に失敗しました')
      console.error('useGroupDetail error:', err)
    } finally {
      setIsLoading(false)
    }
  }

  const refetch = async () => {
    await fetchGroupDetail()
  }

  useEffect(() => {
    if (groupId) {
      fetchGroupDetail()
    }
  }, [groupId])

  return {
    data,
    isLoading,
    error,
    refetch,
  }
}
