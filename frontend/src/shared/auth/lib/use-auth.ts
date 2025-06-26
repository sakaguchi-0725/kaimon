import { useEffect, useState } from 'react'
import * as SecureStore from 'expo-secure-store'

export const useAuth = () => {
  const [isAuth, setIsAuth] = useState(false)

  useEffect(() => {
    const checkAuth = async () => {
      const token = await SecureStore.getItemAsync('accessToken')
      setIsAuth(!!token)
    }
    checkAuth()
  }, [])

  return { isAuth }
}
