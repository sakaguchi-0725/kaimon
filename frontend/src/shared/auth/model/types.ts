import { createContext } from 'react'
import { ApiSchema } from '@/shared/api'

export type GetAccountResponse = ApiSchema<'GetAccountResponse'>

export interface AccountContextType {
  account: GetAccountResponse | undefined
  isLoading: boolean
  error: string | undefined
  refetch: () => void
}

export const AccountContext = createContext<AccountContextType | undefined>(
  undefined,
)
