import { ShoppingItemStatus } from './index'
import type { ComponentType } from 'react'

export const STATUS_LABELS: Record<ShoppingItemStatus, string> = {
  "UNPURCHASED": "未購入",
  "IN_CART": "カート内",
  "PURCHASED": "購入済み"
}

// タブ情報の型定義
export interface TabInfo {
  name: string;
  component: ComponentType<any>;
  options: {
    tabBarLabel: string;
  };
}

// タブ名の定義
export const TAB_NAMES = {
  UNPURCHASED: 'Unpurchased',
  IN_CART: 'InCart',
  PURCHASED: 'Purchased'
} 