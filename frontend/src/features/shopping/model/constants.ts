import { ShoppingItemStatus, Category } from './index'
import type { ComponentType } from 'react'

export const STATUS_LABELS: Record<ShoppingItemStatus, string> = {
  UNPURCHASED: '未購入',
  IN_CART: 'カート内',
  PURCHASED: '購入済み',
}

// タブ情報の型定義
export interface TabInfo {
  name: string
  component: ComponentType<any>
  options: {
    tabBarLabel: string
  }
}

// タブ名の定義
export const TAB_NAMES = {
  UNPURCHASED: 'Unpurchased',
  IN_CART: 'InCart',
  PURCHASED: 'Purchased',
}

// ダミーカテゴリー
export const CATEGORIES: Category[] = [
  { id: 1, name: '野菜・果物' },
  { id: 2, name: '肉・魚' },
  { id: 3, name: '乳製品・卵' },
  { id: 4, name: '飲料' },
  { id: 5, name: '調味料' },
  { id: 6, name: '加工食品' },
  { id: 7, name: '日用品' },
  { id: 8, name: 'その他' },
]
