import { View, StyleSheet, ActivityIndicator, TouchableOpacity } from "react-native"
import { Colors } from "@/shared/constants"
import { FilterModal, SearchBar, FilterChips, STATUS_LABELS, useShoppingList, ShoppingItemList, AddItemModal } from "@/features/shopping"
import { useModal } from "@/shared/ui/modal"
import { Plus } from "react-native-feather"

export const ShoppingItemListScreen = () => {
  const { 
    searchText, 
    setSearchText, 
    selectedFilters, 
    setSelectedFilters,
    removeFilter,
    clearAllFilters,
    hasActiveFilters,
    items,
    isLoading,
    updateItemStatus
  } = useShoppingList()
  
  const { isVisible: filterModalVisible, openModal, closeModal } = useModal()
  const { isVisible: addItemModalVisible, openModal: openAddItemModal, closeModal: closeAddItemModal } = useModal()

  // ダミーグループデータ
  const groups = [
    { id: 'group1', name: '我が家' },
    { id: 'group2', name: '実家' },
    { id: 'group3', name: '職場' }
  ]

  // アイテム追加処理
  const handleAddItem = (name: string, description: string, categoryId: number, groupId: string) => {
    console.log('アイテム追加:', name, description, categoryId, groupId)
    // TODO: アイテム追加のAPI呼び出し
  }

  return (
    <View style={styles.container}>
      {/* 検索バーとフィルターボタン */}
      <SearchBar
        searchText={searchText}
        onSearchTextChange={setSearchText}
        onFilterPress={openModal}
        hasActiveFilters={hasActiveFilters()}
      />

      {/* 選択されたフィルターのチップ表示 */}
      <FilterChips
        selectedFilters={selectedFilters}
        onRemoveFilter={removeFilter}
        onClearAll={clearAllFilters}
        statusLabels={STATUS_LABELS}
      />

      {/* 買い物リスト */}
      {isLoading ? (
        <View style={styles.loadingContainer}>
          <ActivityIndicator size="large" color={Colors.primary} />
        </View>
      ) : (
        <ShoppingItemList
          items={items}
          onStatusChange={updateItemStatus}
        />
      )}

      {/* 新規メモ追加ボタン */}
      <TouchableOpacity
        style={styles.addButton}
        onPress={openAddItemModal}
        activeOpacity={0.8}
      >
        <Plus width={24} height={24} stroke={Colors.white} />
      </TouchableOpacity>

      {/* フィルターモーダル */}
      <FilterModal 
        isVisible={filterModalVisible}
        onClose={closeModal}
        selectedFilters={selectedFilters}
        onFilterChange={setSelectedFilters}
      />

      {/* アイテム追加モーダル */}
      <AddItemModal
        isVisible={addItemModalVisible}
        onClose={closeAddItemModal}
        onAddItem={handleAddItem}
        groups={groups}
      />
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: Colors.backgroundGray,
  },
  loadingContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  addButton: {
    position: 'absolute',
    bottom: 24,
    right: 24,
    width: 56,
    height: 56,
    borderRadius: 28,
    backgroundColor: Colors.primary,
    justifyContent: 'center',
    alignItems: 'center',
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 1 },
    shadowOpacity: 0.2,
    shadowRadius: 2,
  },
})