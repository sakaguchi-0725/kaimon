import {
  createUserWithEmailAndPassword,
  signInWithEmailAndPassword,
  signOut,
} from 'firebase/auth'
import { auth } from '@/shared/lib/firebase'

export const authService = {
  async signUp(email: string, password: string) {
    try {
      const userCredential = await createUserWithEmailAndPassword(
        auth,
        email,
        password,
      )
      return userCredential.user
    } catch (error) {
      throw error
    }
  },

  async signIn(email: string, password: string) {
    try {
      const userCredential = await signInWithEmailAndPassword(
        auth,
        email,
        password,
      )
      return userCredential.user
    } catch (error) {
      throw error
    }
  },

  async signOut() {
    try {
      await signOut(auth)
    } catch (error) {
      throw error
    }
  },
}
