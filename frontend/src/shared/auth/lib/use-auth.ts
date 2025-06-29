import { useEffect, useState } from 'react'
import { getAuth, onAuthStateChanged } from '@react-native-firebase/auth'

export const useAuth = () => {
  const [isAuth, setIsAuth] = useState(false)

  useEffect(() => {
    const auth = getAuth()
    const unsubscribe = onAuthStateChanged(auth, (user) => {
      setIsAuth(!!user)
    })

    return unsubscribe
  }, [])

  return { isAuth }
}
