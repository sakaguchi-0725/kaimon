export type ShoppingItemStatus = 'UNPURCHASED' | 'IN_CART' | 'PURCHASED'

export type ShoppingItem = {
  id: number
  name: string
  description: string
  status: ShoppingItemStatus
  memberId: string
  categoryId: number
}

// 定数のエクスポート
export * from './constants'