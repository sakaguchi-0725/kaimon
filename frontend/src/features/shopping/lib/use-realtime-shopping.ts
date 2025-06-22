import { useState, useEffect } from 'react'
import { ShoppingItem, ShoppingItemStatus } from '../model'

// ダミーデータ（後で削除）
const dummyItems: ShoppingItem[] = [
  { id: 1, name: '牛乳', description: '1L', status: 'UNPURCHASED', memberId: '1', categoryId: 1 },
  { id: 2, name: 'パン', description: '食パン 6枚切り', status: 'UNPURCHASED', memberId: '2', categoryId: 2 },
  { id: 3, name: 'バナナ', description: '3本', status: 'UNPURCHASED', memberId: '1', categoryId: 3 },
  { id: 4, name: '卵', description: '10個入り', status: 'UNPURCHASED', memberId: '3', categoryId: 1 },
  { id: 5, name: 'トマト', description: '3個', status: 'IN_CART', memberId: '1', categoryId: 3 },
  { id: 6, name: 'じゃがいも', description: '5個', status: 'IN_CART', memberId: '2', categoryId: 3 },
  { id: 7, name: '豚肉', description: '300g', status: 'IN_CART', memberId: '3', categoryId: 4 },
  { id: 8, name: 'りんご', description: '2個', status: 'PURCHASED', memberId: '1', categoryId: 3 },
  { id: 9, name: '鶏肉', description: '200g', status: 'PURCHASED', memberId: '2', categoryId: 4 },
]

export interface UseRealtimeShoppingProps {
  groupId?: string
}

export const useRealtimeShopping = ({ groupId }: UseRealtimeShoppingProps = {}) => {
  const [items, setItems] = useState<ShoppingItem[]>(dummyItems)
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  // グループIDが変更されたときにアイテムを再取得
  useEffect(() => {
    if (!groupId) return

    const fetchItems = async () => {
      setIsLoading(true)
      setError(null)
      
      try {
        // 実際のAPIリクエストはここで行う
        // const response = await api.getShoppingItems(groupId)
        // setItems(response.data)
        
        // ダミーデータを使用（遅延をシミュレート）
        setTimeout(() => {
          setItems(dummyItems)
          setIsLoading(false)
        }, 500)
      } catch (err) {
        setError('買い物リストの取得に失敗しました')
        setIsLoading(false)
      }
    }

    fetchItems()
  }, [groupId])

  // ステータス別にアイテムをフィルタリング
  const unpurchasedItems = items.filter(item => item.status === 'UNPURCHASED')
  const inCartItems = items.filter(item => item.status === 'IN_CART')
  const purchasedItems = items.filter(item => item.status === 'PURCHASED')

  // アイテムのステータスを更新する関数
  const updateItemStatus = (itemId: number, newStatus: ShoppingItemStatus) => {
    setItems(prevItems => 
      prevItems.map(item => 
        item.id === itemId ? { ...item, status: newStatus } : item
      )
    )
    
    // 実際のAPIリクエストはここで行う
    // api.updateItemStatus(itemId, newStatus)
  }

  // カートに追加
  const addToCart = (item: ShoppingItem) => {
    updateItemStatus(item.id, 'IN_CART')
  }

  // 購入済みにする
  const markAsPurchased = (item: ShoppingItem) => {
    updateItemStatus(item.id, 'PURCHASED')
  }

  // 未購入に戻す
  const returnToUnpurchased = (item: ShoppingItem) => {
    updateItemStatus(item.id, 'UNPURCHASED')
  }

  // カートに戻す
  const returnToCart = (item: ShoppingItem) => {
    updateItemStatus(item.id, 'IN_CART')
  }

  return {
    isLoading,
    error,
    unpurchasedItems,
    inCartItems,
    purchasedItems,
    addToCart,
    markAsPurchased,
    returnToUnpurchased,
    returnToCart
  }
} 