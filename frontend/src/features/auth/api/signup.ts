import { createUserWithEmailAndPassword } from 'firebase/auth'
import { auth } from '@/shared/lib/firebase'

export const signUp = async (email: string, password: string) => {
  const userCredential = await createUserWithEmailAndPassword(
    auth,
    email,
    password,
  )
  return userCredential.user
}
