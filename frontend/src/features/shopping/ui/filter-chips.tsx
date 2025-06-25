import {
  ScrollView,
  StyleSheet,
  Text,
  TouchableOpacity,
  View,
} from 'react-native'
import { XCircle } from 'react-native-feather'
import { Colors } from '@/shared/constants'
import { ShoppingItemStatus } from '@/features/shopping/model'

export interface FilterChipsProps {
  selectedFilters: {
    categories: string[]
    groups: string[]
    statuses: ShoppingItemStatus[]
  }
  onRemoveFilter: (
    type: 'categories' | 'groups' | 'statuses',
    value: string | ShoppingItemStatus,
  ) => void
  onClearAll: () => void
  statusLabels: Record<ShoppingItemStatus, string>
}

export const FilterChips = ({
  selectedFilters,
  onRemoveFilter,
  onClearAll,
  statusLabels,
}: FilterChipsProps) => {
  const hasActiveFilters = () => {
    return (
      selectedFilters.categories.length > 0 ||
      selectedFilters.groups.length > 0 ||
      selectedFilters.statuses.length > 0
    )
  }

  if (!hasActiveFilters()) {
    return null
  }

  return (
    <ScrollView
      horizontal
      showsHorizontalScrollIndicator={false}
      style={styles.filtersScrollView}
      contentContainerStyle={styles.filtersContainer}
    >
      {selectedFilters.categories.map((category) => (
        <TouchableOpacity
          key={`category-${category}`}
          style={styles.filterChip}
          onPress={() => onRemoveFilter('categories', category)}
        >
          <Text style={styles.filterChipText}>{category}</Text>
          <XCircle width={16} height={16} stroke={Colors.subText} />
        </TouchableOpacity>
      ))}
      {selectedFilters.groups.map((group) => (
        <TouchableOpacity
          key={`group-${group}`}
          style={styles.filterChip}
          onPress={() => onRemoveFilter('groups', group)}
        >
          <Text style={styles.filterChipText}>{group}</Text>
          <XCircle width={16} height={16} stroke={Colors.subText} />
        </TouchableOpacity>
      ))}
      {selectedFilters.statuses.map((status) => (
        <TouchableOpacity
          key={`status-${status}`}
          style={styles.filterChip}
          onPress={() => onRemoveFilter('statuses', status)}
        >
          <Text style={styles.filterChipText}>{statusLabels[status]}</Text>
          <XCircle width={16} height={16} stroke={Colors.subText} />
        </TouchableOpacity>
      ))}
      <TouchableOpacity style={styles.clearAllChip} onPress={onClearAll}>
        <Text style={styles.clearAllText}>すべて解除</Text>
      </TouchableOpacity>
    </ScrollView>
  )
}

const styles = StyleSheet.create({
  filtersScrollView: {
    maxHeight: 50,
    backgroundColor: Colors.white,
  },
  filtersContainer: {
    paddingHorizontal: 16,
    paddingVertical: 8,
    flexDirection: 'row',
  },
  filterChip: {
    flexDirection: 'row',
    alignItems: 'center',
    backgroundColor: Colors.backgroundGray,
    borderRadius: 16,
    paddingHorizontal: 12,
    paddingVertical: 6,
    marginRight: 8,
  },
  filterChipText: {
    fontSize: 14,
    color: Colors.mainText,
    marginRight: 4,
  },
  clearAllChip: {
    flexDirection: 'row',
    alignItems: 'center',
    borderWidth: 1,
    borderColor: Colors.border,
    borderRadius: 16,
    paddingHorizontal: 12,
    paddingVertical: 6,
  },
  clearAllText: {
    fontSize: 14,
    color: Colors.subText,
  },
})
