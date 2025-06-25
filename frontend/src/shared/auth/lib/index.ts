import { useEffect, useState } from 'react'
import { authStorage } from '@/shared/lib/auth-storage'

export const useAuth = () => {
  const [isAuth, setIsAuth] = useState(false)
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    const checkAuth = async () => {
      try {
        const accessToken = await authStorage.getAccessToken()
        setIsAuth(!!accessToken)
      } catch (error) {
        console.error('認証状態の確認でエラーが発生しました:', error)
        setIsAuth(false)
      } finally {
        setIsLoading(false)
      }
    }
    checkAuth()
  }, [])

  const login = async (accessToken: string, refreshToken: string) => {
    await authStorage.setTokens(accessToken, refreshToken)
    setIsAuth(true)
  }

  const logout = async () => {
    await authStorage.clearTokens()
    setIsAuth(false)
  }

  const refreshAuth = async () => {
    const accessToken = await authStorage.getAccessToken()
    setIsAuth(!!accessToken)
  }

  return {
    isAuth,
    isLoading,
    login,
    logout,
    refreshAuth,
  }
}
