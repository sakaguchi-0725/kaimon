const SESSION_KEY = 'email'

export const setSession = (email: string) => {
  sessionStorage.setItem(SESSION_KEY, email)
}

export const getSession = () => {
  return sessionStorage.getItem(SESSION_KEY)
}

export const removeSession = () => {
  sessionStorage.removeItem(SESSION_KEY)
}
