export type ToastType = 'success' | 'error'

export type ToastItem = {
  id: number
  type: ToastType
  title: string
  message: string
  removing?: boolean
}

export type ToastContainerRef = {
  showError: (title: string | undefined, message: string) => void
  showSuccess: (title: string | undefined, message: string) => void
}
