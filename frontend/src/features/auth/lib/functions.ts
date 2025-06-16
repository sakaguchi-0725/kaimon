const SESSION_KEY = 'email'

export const setEmail = (email: string) => {
  sessionStorage.setItem(SESSION_KEY, email)
}

export const getEmail = () => {
  return sessionStorage.getItem(SESSION_KEY)
}

export const removeEmail = () => {
  sessionStorage.removeItem(SESSION_KEY)
}
