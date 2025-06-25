import { signOut as firebaseSignOut } from 'firebase/auth'
import { auth } from '@/shared/lib/firebase'

export const signOut = async () => {
  await firebaseSignOut(auth)
}
