import { View, Text, StyleSheet, FlatList } from 'react-native'
import { Colors } from '@/shared/constants'
import { ShoppingItem } from '../model'
import { ShoppingItemCard } from './shopping-item-card'
import { ActionButton } from './action-button'
import { ShoppingCart } from 'react-native-feather'

// ダミーデータ（後で削除）
const dummyItems: ShoppingItem[] = [
  {
    id: 1,
    name: '牛乳',
    description: '1L',
    status: 'UNPURCHASED',
    memberId: '1',
    categoryId: 1,
  },
  {
    id: 2,
    name: 'パン',
    description: '食パン 6枚切り',
    status: 'UNPURCHASED',
    memberId: '2',
    categoryId: 2,
  },
  {
    id: 3,
    name: 'バナナ',
    description: '3本',
    status: 'UNPURCHASED',
    memberId: '1',
    categoryId: 3,
  },
  {
    id: 4,
    name: '卵',
    description: '10個入り',
    status: 'UNPURCHASED',
    memberId: '3',
    categoryId: 1,
  },
]

export interface UnpurchasedItemsTabProps {
  items?: ShoppingItem[]
  onAddToCart?: (item: ShoppingItem) => void
}

export const UnpurchasedItemsTab = ({
  items = dummyItems, // 本来はpropsから受け取る
  onAddToCart,
}: UnpurchasedItemsTabProps) => {
  // アイテムがない場合は空の表示
  if (items.length === 0) {
    return (
      <View style={styles.emptyContainer}>
        <Text style={styles.emptyText}>未購入のアイテムがありません</Text>
      </View>
    )
  }

  return (
    <FlatList
      data={items}
      keyExtractor={(item) => item.id.toString()}
      renderItem={({ item }) => (
        <ShoppingItemCard
          item={item}
          actionButtons={
            <ActionButton
              icon={
                <ShoppingCart width={20} height={20} stroke={Colors.white} />
              }
              label="カートに入れる"
              status="UNPURCHASED"
              kind="add"
              onPress={() => onAddToCart?.(item)}
            />
          }
        />
      )}
      contentContainerStyle={styles.listContent}
      showsVerticalScrollIndicator={false}
    />
  )
}

const styles = StyleSheet.create({
  emptyContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: Colors.backgroundGray,
  },
  emptyText: {
    fontSize: 16,
    color: Colors.subText,
  },
  listContent: {
    paddingVertical: 8,
    backgroundColor: Colors.backgroundGray,
    flexGrow: 1,
  },
})
