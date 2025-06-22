import { Text, View, StyleSheet } from "react-native"
import { Colors } from "@/shared/constants"
import { FilterModal, SearchBar, FilterChips, STATUS_LABELS, useShoppingList } from "@/features/shopping"
import { useModal } from "@/shared/ui/modal"

export const ShoppingItemListScreen = () => {
  const { 
    searchText, 
    setSearchText, 
    selectedFilters, 
    setSelectedFilters,
    removeFilter,
    clearAllFilters,
    hasActiveFilters 
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

      {/* ここに買い物リストを表示 */}
      <View style={styles.listContainer}>
        <Text style={styles.emptyText}>買い物リストがここに表示されます</Text>
      </View>

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
  listContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  emptyText: {
    fontSize: 16,
    color: Colors.subText,
  },
})