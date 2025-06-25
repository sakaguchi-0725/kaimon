import { useEffect, useState } from 'react'
import { ShoppingItem, ShoppingItemStatus } from '../model'

export interface ShoppingListFilters {
  categories: string[]
  groups: string[]
  statuses: ShoppingItemStatus[]
}

// ダミーデータ
const DUMMY_ITEMS: ShoppingItem[] = [
  {
    id: 1,
    name: '牛乳',
    description: '1リットルの紙パック',
    status: 'UNPURCHASED',
    memberId: 'user1',
    categoryId: 1,
  },
  {
    id: 2,
    name: '食パン',
    description: '6枚切り',
    status: 'IN_CART',
    memberId: 'user2',
    categoryId: 1,
  },
  {
    id: 3,
    name: 'トイレットペーパー',
    description: '12ロール入り',
    status: 'PURCHASED',
    memberId: 'user1',
    categoryId: 2,
  },
  {
    id: 4,
    name: 'シャンプー',
    description: 'ボトル500ml',
    status: 'UNPURCHASED',
    memberId: 'user3',
    categoryId: 2,
  },
  {
    id: 5,
    name: 'バナナ',
    description: '1房',
    status: 'UNPURCHASED',
    memberId: 'user2',
    categoryId: 1,
  },
]

export const useShoppingList = () => {
  // 検索テキスト
  const [searchText, setSearchText] = useState('')

  // フィルター条件
  const [selectedFilters, setSelectedFilters] = useState<ShoppingListFilters>({
    categories: [],
    groups: [],
    statuses: [],
  })

  // 買い物アイテムリスト
  const [items, setItems] = useState<ShoppingItem[]>([])
  const [filteredItems, setFilteredItems] = useState<ShoppingItem[]>([])
  const [isLoading, setIsLoading] = useState(false)

  // 初期データの読み込み
  useEffect(() => {
    const loadItems = async () => {
      setIsLoading(true)
      try {
        // 実際の実装では、APIからデータを取得する
        // const response = await fetch('/api/shopping-items')
        // const data = await response.json()
        // setItems(data)

        // ダミーデータを使用
        setItems(DUMMY_ITEMS)
      } catch (error) {
        console.error('Failed to load shopping items:', error)
      } finally {
        setIsLoading(false)
      }
    }

    loadItems()
  }, [])

  // フィルターと検索テキストが変更されたときにアイテムをフィルタリング
  useEffect(() => {
    let result = [...items]

    // フィルターの適用
    if (selectedFilters.categories.length > 0) {
      result = result.filter((item) =>
        selectedFilters.categories.includes(item.categoryId.toString()),
      )
    }

    if (selectedFilters.groups.length > 0) {
      // 実際の実装では、グループによるフィルタリングを行う
      // 現在はダミーデータなのでスキップ
    }

    if (selectedFilters.statuses.length > 0) {
      result = result.filter((item) =>
        selectedFilters.statuses.includes(item.status),
      )
    }

    // 検索テキストの適用
    if (searchText) {
      const lowerSearchText = searchText.toLowerCase()
      result = result.filter(
        (item) =>
          item.name.toLowerCase().includes(lowerSearchText) ||
          (item.description &&
            item.description.toLowerCase().includes(lowerSearchText)),
      )
    }

    setFilteredItems(result)
  }, [items, selectedFilters, searchText])

  /**
   * フィルターを削除する
   * @param type フィルターの種類
   * @param value 削除する値
   */
  const removeFilter = (
    type: keyof ShoppingListFilters,
    value: string | ShoppingItemStatus,
  ) => {
    setSelectedFilters((prev) => {
      const currentFilters = [...prev[type]]
      const index = currentFilters.indexOf(value as never)

      if (index >= 0) {
        currentFilters.splice(index, 1)
      }

      return { ...prev, [type]: currentFilters }
    })
  }

  const clearAllFilters = () => {
    setSelectedFilters({ categories: [], groups: [], statuses: [] })
  }

  /**
   * アクティブなフィルターがあるかどうかを判定する
   * @returns アクティブなフィルターがある場合はtrue
   */
  const hasActiveFilters = () => {
    return (
      selectedFilters.categories.length > 0 ||
      selectedFilters.groups.length > 0 ||
      selectedFilters.statuses.length > 0
    )
  }

  /**
   * アイテムのステータスを更新する
   * @param item 更新するアイテム
   * @param newStatus 新しいステータス
   */
  const updateItemStatus = (
    item: ShoppingItem,
    newStatus: ShoppingItemStatus,
  ) => {
    // 実際の実装では、APIを呼び出してステータスを更新する
    // const updateItem = async () => {
    //   try {
    //     const response = await fetch(`/api/shopping-items/${item.id}`, {
    //       method: 'PATCH',
    //       headers: {
    //         'Content-Type': 'application/json',
    //       },
    //       body: JSON.stringify({ status: newStatus }),
    //     })
    //     if (!response.ok) throw new Error('Failed to update item status')
    //   } catch (error) {
    //     console.error('Error updating item status:', error)
    //   }
    // }
    // updateItem()

    // ローカルの状態を更新
    setItems((prevItems) =>
      prevItems.map((prevItem) =>
        prevItem.id === item.id ? { ...prevItem, status: newStatus } : prevItem,
      ),
    )
  }

  return {
    searchText,
    setSearchText,
    selectedFilters,
    setSelectedFilters,
    removeFilter,
    clearAllFilters,
    hasActiveFilters,
    items: filteredItems,
    isLoading,
    updateItemStatus,
  }
}
