import { useContext } from 'react'
import { AccountContext, AccountContextType } from '../model/types'

export const useAccount = (): AccountContextType => {
  const context = useContext(AccountContext)

  if (context === undefined) {
    throw new Error('useAccount must be used within an AccountContextProvider')
  }

  return context
}
