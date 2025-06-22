export type ShoppingItemStatus = 'UNPURCHASED' | 'IN_CART' | 'PURCHASED'

export type ShoppingItem = {
  id: number
  name: string
  description: string
  status: ShoppingItemStatus
  memberId: string
  categoryId: number
}

// カテゴリー情報の定義
export interface Category {
  id: number;
  name: string;
}

// 定数のエクスポート
export * from './constants'