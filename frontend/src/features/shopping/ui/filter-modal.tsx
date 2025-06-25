import { useState } from 'react'
import {
  Text,
  View,
  StyleSheet,
  TouchableOpacity,
  ScrollView,
} from 'react-native'
import { X } from 'react-native-feather'
import { Colors } from '@/shared/constants'
import { BottomSheetModal } from '@/shared/ui/modal'
import { Button } from '@/shared/ui'
import { ShoppingItemStatus, STATUS_LABELS } from '@/features/shopping/model'

export interface FilterModalProps {
  isVisible: boolean
  onClose: () => void
  selectedFilters: {
    categories: string[]
    groups: string[]
    statuses: ShoppingItemStatus[]
  }
  onFilterChange: (filters: {
    categories: string[]
    groups: string[]
    statuses: ShoppingItemStatus[]
  }) => void
}

export const FilterModal = ({
  isVisible,
  onClose,
  selectedFilters,
  onFilterChange,
}: FilterModalProps) => {
  // ローカルステートでフィルター選択を管理
  const [localFilters, setLocalFilters] = useState(selectedFilters)

  // ダミーデータ（実際の実装時には置き換える）
  const dummyCategories = ['食品', '日用品', '衣類', 'その他']
  const dummyGroups = ['家族', '友人', '仕事']
  const dummyStatuses: ShoppingItemStatus[] = [
    'UNPURCHASED',
    'IN_CART',
    'PURCHASED',
  ]

  // フィルターの切り替え
  const toggleFilter = (
    type: 'categories' | 'groups' | 'statuses',
    value: string | ShoppingItemStatus,
  ) => {
    setLocalFilters((prev) => {
      const currentFilters = [...prev[type]]
      const index = currentFilters.indexOf(value as never)

      if (index >= 0) {
        currentFilters.splice(index, 1)
      } else {
        currentFilters.push(value as never)
      }

      return { ...prev, [type]: currentFilters }
    })
  }

  // フィルターを適用して閉じる
  const applyFilters = () => {
    onFilterChange(localFilters)
    onClose()
  }

  return (
    <BottomSheetModal isVisible={isVisible} onClose={onClose}>
      <View style={styles.modalHeader}>
        <Text style={styles.modalTitle}>絞り込み</Text>
        <TouchableOpacity onPress={onClose}>
          <X width={24} height={24} stroke={Colors.mainText} />
        </TouchableOpacity>
      </View>

      <ScrollView style={styles.scrollView}>
        {/* カテゴリーフィルター */}
        <View style={styles.filterSection}>
          <Text style={styles.filterSectionTitle}>カテゴリー</Text>
          <View style={styles.filterOptions}>
            {dummyCategories.map((category) => (
              <TouchableOpacity
                key={`filter-category-${category}`}
                style={[
                  styles.filterOption,
                  localFilters.categories.includes(category) &&
                    styles.selectedFilterOption,
                ]}
                onPress={() => toggleFilter('categories', category)}
              >
                <Text
                  style={[
                    styles.filterOptionText,
                    localFilters.categories.includes(category) &&
                      styles.selectedFilterOptionText,
                  ]}
                >
                  {category}
                </Text>
              </TouchableOpacity>
            ))}
          </View>
        </View>

        {/* グループフィルター */}
        <View style={styles.filterSection}>
          <Text style={styles.filterSectionTitle}>グループ</Text>
          <View style={styles.filterOptions}>
            {dummyGroups.map((group) => (
              <TouchableOpacity
                key={`filter-group-${group}`}
                style={[
                  styles.filterOption,
                  localFilters.groups.includes(group) &&
                    styles.selectedFilterOption,
                ]}
                onPress={() => toggleFilter('groups', group)}
              >
                <Text
                  style={[
                    styles.filterOptionText,
                    localFilters.groups.includes(group) &&
                      styles.selectedFilterOptionText,
                  ]}
                >
                  {group}
                </Text>
              </TouchableOpacity>
            ))}
          </View>
        </View>

        {/* ステータスフィルター */}
        <View style={styles.filterSection}>
          <Text style={styles.filterSectionTitle}>ステータス</Text>
          <View style={styles.filterOptions}>
            {dummyStatuses.map((status) => (
              <TouchableOpacity
                key={`filter-status-${status}`}
                style={[
                  styles.filterOption,
                  localFilters.statuses.includes(status) &&
                    styles.selectedFilterOption,
                ]}
                onPress={() => toggleFilter('statuses', status)}
              >
                <Text
                  style={[
                    styles.filterOptionText,
                    localFilters.statuses.includes(status) &&
                      styles.selectedFilterOptionText,
                  ]}
                >
                  {STATUS_LABELS[status]}
                </Text>
              </TouchableOpacity>
            ))}
          </View>
        </View>
      </ScrollView>

      {/* フィルター適用ボタン */}
      <Button
        text="適用する"
        onPress={applyFilters}
        size="full"
        variant="solid"
        color="primary"
      />
    </BottomSheetModal>
  )
}

const styles = StyleSheet.create({
  modalHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 16,
  },
  modalTitle: {
    fontSize: 18,
    fontWeight: 'bold',
    color: Colors.mainText,
  },
  scrollView: {
    marginBottom: 8,
  },
  filterSection: {
    marginBottom: 20,
  },
  filterSectionTitle: {
    fontSize: 16,
    fontWeight: 'bold',
    color: Colors.mainText,
    marginBottom: 10,
  },
  filterOptions: {
    flexDirection: 'row',
    flexWrap: 'wrap',
  },
  filterOption: {
    borderWidth: 1,
    borderColor: Colors.border,
    borderRadius: 20,
    paddingHorizontal: 14,
    paddingVertical: 8,
    marginRight: 8,
    marginBottom: 8,
  },
  selectedFilterOption: {
    backgroundColor: Colors.primary,
    borderColor: Colors.primary,
  },
  filterOptionText: {
    fontSize: 14,
    color: Colors.mainText,
  },
  selectedFilterOptionText: {
    color: Colors.white,
  },
})
