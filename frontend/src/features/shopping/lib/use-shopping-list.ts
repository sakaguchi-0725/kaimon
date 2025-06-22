import { useState } from 'react'
import { ShoppingItemStatus } from '../model'

export interface ShoppingListFilters {
  categories: string[]
  groups: string[]
  statuses: ShoppingItemStatus[]
}

export const useShoppingList = () => {
  // 検索テキスト
  const [searchText, setSearchText] = useState("")
  
  // フィルター条件
  const [selectedFilters, setSelectedFilters] = useState<ShoppingListFilters>({
    categories: [],
    groups: [],
    statuses: [],
  })

  /**
   * フィルターを削除する
   * @param type フィルターの種類
   * @param value 削除する値
   */
  const removeFilter = (type: keyof ShoppingListFilters, value: string | ShoppingItemStatus) => {
    setSelectedFilters(prev => {
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
    return selectedFilters.categories.length > 0 || 
           selectedFilters.groups.length > 0 || 
           selectedFilters.statuses.length > 0
  }

  /**
   * フィルターを適用して買い物アイテムを取得する
   * 現在はダミー実装
   */
  const fetchFilteredItems = async () => {
    // 実際の実装では、APIリクエストを送信してフィルター条件に合致するアイテムを取得する
    // const params = new URLSearchParams()
    // if (selectedFilters.categories.length > 0) {
    //   selectedFilters.categories.forEach(cat => params.append('category', cat))
    // }
    // if (selectedFilters.groups.length > 0) {
    //   selectedFilters.groups.forEach(group => params.append('group', group))
    // }
    // if (selectedFilters.statuses.length > 0) {
    //   selectedFilters.statuses.forEach(status => params.append('status', status))
    // }
    // 
    // const response = await fetch(`/api/shopping-items?${params.toString()}`)
    // const data = await response.json()
    // return data
  }

  /**
   * 検索テキストでアイテムをフィルタリングする
   * 現在はダミー実装
   */
  const filterItemsBySearchText = (items: any[]) => {
    // 実際の実装では、検索テキストに合致するアイテムをフィルタリングする
    // return items.filter(item => 
    //   item.name.toLowerCase().includes(searchText.toLowerCase()) || 
    //   item.description.toLowerCase().includes(searchText.toLowerCase())
    // )
    return items
  }

  return {
    searchText,
    setSearchText,
    selectedFilters,
    setSelectedFilters,
    removeFilter,
    clearAllFilters,
    hasActiveFilters,
    fetchFilteredItems,
    filterItemsBySearchText,
  }
} 