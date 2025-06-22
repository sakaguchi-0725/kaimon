import { Text, View, StyleSheet, ActivityIndicator } from "react-native"
import { Colors } from "@/shared/constants"
import { FilterModal, SearchBar, FilterChips, STATUS_LABELS, useShoppingList, ShoppingItemList } from "@/features/shopping"
import { useModal } from "@/shared/ui/modal"

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

      {/* フィルターモーダル */}
      <FilterModal 
        isVisible={filterModalVisible}
        onClose={closeModal}
        selectedFilters={selectedFilters}
        onFilterChange={setSelectedFilters}
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
})