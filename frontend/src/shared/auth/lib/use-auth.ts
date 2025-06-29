import { useEffect, useState } from 'react'
import auth from '@react-native-firebase/auth'

export const useAuth = () => {
  const [isAuth, setIsAuth] = useState(false)

  useEffect(() => {
    const unsubscribe = auth().onAuthStateChanged((user) => {
      setIsAuth(!!user)
    })

    return unsubscribe
  }, [])

  return { isAuth }
}
