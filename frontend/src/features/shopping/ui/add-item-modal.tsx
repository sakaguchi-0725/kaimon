import { useState } from 'react'
import { View, Text, StyleSheet, TouchableOpacity } from 'react-native'
import { Colors } from '@/shared/constants'
import { BottomSheetModal } from '@/shared/ui/modal'
import { Category } from '../model'
import { CATEGORIES } from '../model/constants'
import { X } from 'react-native-feather'
import { Dropdown, DropdownOption } from '@/shared/ui/dropdown'
import { TextInput, Textarea } from '@/shared/ui/input'

type Props = {
  isVisible: boolean
  onClose: () => void
  onAddItem?: (name: string, description: string, categoryId: number, groupId: string) => void
  groups?: { id: string, name: string }[]
}

export const AddItemModal = ({ isVisible, onClose, onAddItem, groups = [] }: Props) => {
  const [itemName, setItemName] = useState('')
  const [itemDescription, setItemDescription] = useState('')
  const [selectedCategory, setSelectedCategory] = useState<Category | null>(null)
  const [selectedGroup, setSelectedGroup] = useState<{ id: string, name: string } | null>(
    groups.length > 0 ? groups[0] : null
  )

  const handleAddItem = () => {
    if (!itemName.trim() || !selectedCategory || !selectedGroup) return

    // アイテム追加処理
    onAddItem?.(itemName, itemDescription, selectedCategory.id, selectedGroup.id)
    
    // フォームをリセット
    setItemName('')
    setItemDescription('')
    setSelectedCategory(null)
    
    // モーダルを閉じる
    onClose()
  }

  // カテゴリーのオプション
  const categoryOptions: DropdownOption<Category>[] = CATEGORIES.map(category => ({
    label: category.name,
    value: category
  }))

  // グループのオプション
  const groupOptions: DropdownOption<{ id: string, name: string }>[] = groups.map(group => ({
    label: group.name,
    value: group
  }))

  const isAddButtonDisabled = !itemName.trim() || !selectedCategory || !selectedGroup

  return (
    <BottomSheetModal isVisible={isVisible} onClose={onClose}>
      <View style={styles.modalHeader}>
        <Text style={styles.title}>買い物メモを追加</Text>
        <TouchableOpacity onPress={onClose}>
          <X width={24} height={24} stroke={Colors.mainText} />
        </TouchableOpacity>
      </View>
      
      <View style={styles.formContainer}>
        {/* 商品名入力 */}
        <TextInput
          label="商品名"
          required
          value={itemName}
          onChangeText={setItemName}
          placeholder="例: 牛乳"
        />
        
        {/* カテゴリー選択 */}
        <Dropdown
          label="カテゴリー"
          required
          selectedValue={selectedCategory}
          onValueChange={setSelectedCategory}
          options={categoryOptions}
          placeholder="カテゴリーを選択"
        />
        
        {/* グループ選択 */}
        {groups.length > 1 && (
          <Dropdown
            label="グループ"
            required
            selectedValue={selectedGroup}
            onValueChange={setSelectedGroup}
            options={groupOptions}
            placeholder="グループを選択"
          />
        )}
        
        {/* 詳細入力 */}
        <Textarea
          label="詳細"
          value={itemDescription}
          onChangeText={setItemDescription}
          placeholder="例: 1L パック 2つ"
          numberOfLines={3}
        />
      </View>
      
      {/* 追加ボタン */}
      <TouchableOpacity
        style={[
          styles.addButton,
          isAddButtonDisabled && styles.addButtonDisabled
        ]}
        onPress={handleAddItem}
        disabled={isAddButtonDisabled}
      >
        <Text style={[
          styles.addButtonText,
          isAddButtonDisabled && styles.addButtonTextDisabled
        ]}>追加</Text>
      </TouchableOpacity>
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
  title: {
    fontSize: 20,
    fontWeight: 'bold',
    color: Colors.mainText,
  },
  formContainer: {
    gap: 16,
    marginBottom: 8,
  },
  addButton: {
    backgroundColor: Colors.primary,
    paddingVertical: 12,
    borderRadius: 8,
    alignItems: 'center',
    marginTop: 8,
  },
  addButtonDisabled: {
    backgroundColor: Colors.secondary,
  },
  addButtonText: {
    color: Colors.white,
    fontSize: 16,
    fontWeight: '500',
  },
  addButtonTextDisabled: {
    color: Colors.subText,
  },
}) 