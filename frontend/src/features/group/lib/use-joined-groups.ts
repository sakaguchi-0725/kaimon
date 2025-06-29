import { useState, useCallback } from 'react'
import { useFocusEffect } from '@react-navigation/native'
import { client } from '@/shared/api/client'
import type { JoinedGroup } from '../model'

type UseJoinedGroupsReturn = {
  groups: JoinedGroup[]
  isLoading: boolean
  error: string | undefined
}

export const useJoinedGroups = (): UseJoinedGroupsReturn => {
  const [groups, setGroups] = useState<JoinedGroup[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string>()

  const fetchGroups = async () => {
    setIsLoading(true)
    setError(undefined)

    try {
      const { data, error: apiError } = await client.GET('/groups')

      if (apiError) {
        setError(apiError.message || 'グループ一覧の取得に失敗しました')
        return
      }

      if (data?.groups) {
        setGroups(data.groups)
      }
    } catch (err) {
      console.error('Failed to fetch joined groups:', err)
      setError('グループ一覧の取得に失敗しました')
    } finally {
      setIsLoading(false)
    }
  }

  useFocusEffect(
    useCallback(() => {
      let isActive = true

      const fetchData = async () => {
        if (isActive) {
          await fetchGroups()
        }
      }

      fetchData()

      return () => {
        isActive = false
      }
    }, []),
  )

  return {
    groups,
    isLoading,
    error,
  }
}
