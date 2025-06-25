import { View, Text, StyleSheet, FlatList } from 'react-native'
import { Colors } from '@/shared/constants'
import { ShoppingItem } from '../model'
import { ShoppingItemCard } from './shopping-item-card'
import { ActionButton } from './action-button'
import { ShoppingCart } from 'react-native-feather'

// ダミーデータ（後で削除）
const dummyItems: ShoppingItem[] = [
  {
    id: 8,
    name: 'りんご',
    description: '2個',
    status: 'PURCHASED',
    memberId: '1',
    categoryId: 3,
  },
  {
    id: 9,
    name: '鶏肉',
    description: '200g',
    status: 'PURCHASED',
    memberId: '2',
    categoryId: 4,
  },
]

export interface PurchasedItemsTabProps {
  items?: ShoppingItem[]
  onReturnToCart?: (item: ShoppingItem) => void
}

export const PurchasedItemsTab = ({
  items = dummyItems, // 本来はpropsから受け取る
  onReturnToCart,
}: PurchasedItemsTabProps) => {
  // アイテムがない場合は空の表示
  if (items.length === 0) {
    return (
      <View style={styles.emptyContainer}>
        <Text style={styles.emptyText}>購入済みのアイテムがありません</Text>
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
                <ShoppingCart width={20} height={20} stroke={Colors.subText} />
              }
              label="カートに戻す"
              status="PURCHASED"
              kind="return"
              onPress={() => onReturnToCart?.(item)}
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
