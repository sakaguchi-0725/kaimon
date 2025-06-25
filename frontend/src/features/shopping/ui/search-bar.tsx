import { StyleSheet, TextInput, TouchableOpacity, View } from 'react-native'
import { Search, X, MoreVertical } from 'react-native-feather'
import { Colors } from '@/shared/constants'

export interface SearchBarProps {
  searchText: string
  onSearchTextChange: (text: string) => void
  onFilterPress: () => void
  hasActiveFilters: boolean
}

export const SearchBar = ({
  searchText,
  onSearchTextChange,
  onFilterPress,
  hasActiveFilters,
}: SearchBarProps) => {
  return (
    <View style={styles.searchContainer}>
      <View style={styles.searchBar}>
        <Search
          width={20}
          height={20}
          stroke={Colors.subText}
          style={styles.searchIcon}
        />
        <TextInput
          style={styles.searchInput}
          placeholder="買い物アイテムを検索"
          value={searchText}
          onChangeText={onSearchTextChange}
        />
        {searchText.length > 0 && (
          <TouchableOpacity onPress={() => onSearchTextChange('')}>
            <X width={20} height={20} stroke={Colors.subText} />
          </TouchableOpacity>
        )}
      </View>
      <TouchableOpacity
        style={[
          styles.filterButton,
          hasActiveFilters && styles.activeFilterButton,
        ]}
        onPress={onFilterPress}
      >
        <MoreVertical
          width={20}
          height={20}
          stroke={hasActiveFilters ? Colors.white : Colors.subText}
        />
      </TouchableOpacity>
    </View>
  )
}

const styles = StyleSheet.create({
  searchContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    paddingHorizontal: 16,
    paddingVertical: 12,
    backgroundColor: Colors.white,
    borderBottomWidth: 1,
    borderBottomColor: Colors.border,
  },
  searchBar: {
    flex: 1,
    flexDirection: 'row',
    alignItems: 'center',
    backgroundColor: Colors.backgroundGray,
    borderRadius: 8,
    paddingHorizontal: 10,
    height: 40,
  },
  searchIcon: {
    marginRight: 6,
  },
  searchInput: {
    flex: 1,
    height: 40,
    fontSize: 16,
    color: Colors.mainText,
  },
  filterButton: {
    marginLeft: 12,
    width: 40,
    height: 40,
    borderRadius: 8,
    backgroundColor: Colors.backgroundGray,
    justifyContent: 'center',
    alignItems: 'center',
  },
  activeFilterButton: {
    backgroundColor: Colors.primary,
  },
})
