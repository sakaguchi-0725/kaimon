import { ReactNode } from 'react'
import {
  useAuth,
  useAccountApi,
  AccountContext,
  AccountContextType,
} from '@/shared/auth'

type Props = {
  children: ReactNode
}

export const AccountContextProvider = ({ children }: Props) => {
  const { isAuth } = useAuth()

  const { account, isLoading, error, refetch } = useAccountApi(isAuth)

  const contextValue: AccountContextType = {
    account,
    isLoading: isAuth ? isLoading : false,
    error,
    refetch,
  }

  return (
    <AccountContext.Provider value={contextValue}>
      {children}
    </AccountContext.Provider>
  )
}
