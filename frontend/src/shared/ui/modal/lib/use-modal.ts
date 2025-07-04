import { useState } from 'react'

export const useModal = () => {
  const [isVisible, setIsVisible] = useState(false)

  const openModal = () => {
    setIsVisible(true)
  }

  const closeModal = () => {
    setIsVisible(false)
  }

  const toggleModal = () => {
    setIsVisible((prev) => !prev)
  }

  return {
    isVisible,
    openModal,
    closeModal,
    toggleModal,
  }
}
