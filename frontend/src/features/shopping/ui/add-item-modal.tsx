import { useState } from 'react'
import { View, Text, StyleSheet, TextInput, TouchableOpacity, ScrollView } from 'react-native'
import { Colors } from '@/shared/constants'
import { BottomSheetModal } from '@/shared/ui/modal'
import { Category } from '../model'
import { CATEGORIES } from '../model/constants'
import { ChevronDown, X } from 'react-native-feather'

export interface AddItemModalProps {
  isVisible: boolean
  onClose: () => void
  onAddItem?: (name: string, description: string, categoryId: number, groupId: string) => void
  groups?: { id: string, name: string }[]
}

export const AddItemModal = ({ isVisible, onClose, onAddItem, groups = [] }: AddItemModalProps) => {
  const [itemName, setItemName] = useState('')
  const [itemDescription, setItemDescription] = useState('')
  const [selectedCategory, setSelectedCategory] = useState<Category | null>(null)
  const [selectedGroup, setSelectedGroup] = useState<{ id: string, name: string } | null>(
    groups.length > 0 ? groups[0] : null
  )
  const [showCategorySelector, setShowCategorySelector] = useState(false)
  const [showGroupSelector, setShowGroupSelector] = useState(false)

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

  const isAddButtonDisabled = !itemName.trim() || !selectedCategory || !selectedGroup

  return (
    <BottomSheetModal isVisible={isVisible} onClose={onClose}>
      <View style={styles.modalHeader}>
        <Text style={styles.title}>買い物メモを追加</Text>
        <TouchableOpacity onPress={onClose}>
          <X width={24} height={24} stroke={Colors.mainText} />
        </TouchableOpacity>
      </View>
      
      <ScrollView style={styles.scrollView}>
        {/* 商品名入力 */}
        <View style={styles.inputContainer}>
          <Text style={styles.label}>商品名 <Text style={styles.required}>*</Text></Text>
          <TextInput
            style={styles.input}
            value={itemName}
            onChangeText={setItemName}
            placeholder="例: 牛乳"
            placeholderTextColor={Colors.subText + '80'}
          />
        </View>
        
        {/* カテゴリー選択 */}
        <View style={styles.inputContainer}>
          <Text style={styles.label}>カテゴリー <Text style={styles.required}>*</Text></Text>
          <TouchableOpacity 
            style={styles.selector}
            onPress={() => setShowCategorySelector(!showCategorySelector)}
          >
            <Text style={selectedCategory ? styles.selectorText : styles.selectorPlaceholder}>
              {selectedCategory ? selectedCategory.name : 'カテゴリーを選択'}
            </Text>
            <ChevronDown width={20} height={20} stroke={Colors.subText} />
          </TouchableOpacity>
          
          {showCategorySelector && (
            <View style={styles.dropdownContainer}>
              {CATEGORIES.map((category) => (
                <TouchableOpacity
                  key={category.id}
                  style={styles.dropdownItem}
                  onPress={() => {
                    setSelectedCategory(category)
                    setShowCategorySelector(false)
                  }}
                >
                  <Text style={styles.dropdownItemText}>{category.name}</Text>
                </TouchableOpacity>
              ))}
            </View>
          )}
        </View>
        
        {/* グループ選択 */}
        {groups.length > 1 && (
          <View style={styles.inputContainer}>
            <Text style={styles.label}>グループ <Text style={styles.required}>*</Text></Text>
            <TouchableOpacity 
              style={styles.selector}
              onPress={() => setShowGroupSelector(!showGroupSelector)}
            >
              <Text style={selectedGroup ? styles.selectorText : styles.selectorPlaceholder}>
                {selectedGroup ? selectedGroup.name : 'グループを選択'}
              </Text>
              <ChevronDown width={20} height={20} stroke={Colors.subText} />
            </TouchableOpacity>
            
            {showGroupSelector && (
              <View style={styles.dropdownContainer}>
                {groups.map((group) => (
                  <TouchableOpacity
                    key={group.id}
                    style={styles.dropdownItem}
                    onPress={() => {
                      setSelectedGroup(group)
                      setShowGroupSelector(false)
                    }}
                  >
                    <Text style={styles.dropdownItemText}>{group.name}</Text>
                  </TouchableOpacity>
                ))}
              </View>
            )}
          </View>
        )}
        
        {/* 詳細入力 */}
        <View style={styles.inputContainer}>
          <Text style={styles.label}>詳細</Text>
          <TextInput
            style={[styles.input, styles.textArea]}
            value={itemDescription}
            onChangeText={setItemDescription}
            placeholder="例: 1L パック 2つ"
            placeholderTextColor={Colors.subText + '80'}
            multiline
            numberOfLines={3}
          />
        </View>
      </ScrollView>
      
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
  scrollView: {
    marginBottom: 8,
  },
  inputContainer: {
    marginBottom: 16,
  },
  label: {
    fontSize: 14,
    fontWeight: '500',
    color: Colors.mainText,
    marginBottom: 8,
  },
  required: {
    color: Colors.error,
  },
  input: {
    backgroundColor: Colors.backgroundGray,
    borderRadius: 8,
    padding: 12,
    fontSize: 16,
    color: Colors.mainText,
  },
  textArea: {
    height: 80,
    textAlignVertical: 'top',
  },
  selector: {
    backgroundColor: Colors.backgroundGray,
    borderRadius: 8,
    padding: 12,
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
  },
  selectorText: {
    fontSize: 16,
    color: Colors.mainText,
  },
  selectorPlaceholder: {
    fontSize: 16,
    color: Colors.subText + '80',
  },
  dropdownContainer: {
    backgroundColor: Colors.white,
    borderRadius: 8,
    marginTop: 4,
    borderWidth: 1,
    borderColor: Colors.border,
    maxHeight: 200,
  },
  dropdownItem: {
    padding: 12,
    borderBottomWidth: 1,
    borderBottomColor: Colors.border,
  },
  dropdownItemText: {
    fontSize: 16,
    color: Colors.mainText,
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