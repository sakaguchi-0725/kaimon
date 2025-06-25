import { signInWithEmailAndPassword } from 'firebase/auth'
import { auth } from '@/shared/lib/firebase'

export const signIn = async (email: string, password: string) => {
  const userCredential = await signInWithEmailAndPassword(auth, email, password)
  return userCredential.user
}
