import { View, Text, StyleSheet, FlatList } from 'react-native'
import { Colors } from '@/shared/constants'
import { ShoppingItem } from '../model'
import { ShoppingItemCard } from './shopping-item-card'
import { ActionButton, ActionButtonContainer } from './action-button'
import { Check, CornerDownLeft } from 'react-native-feather'

// ダミーデータ（後で削除）
const dummyItems: ShoppingItem[] = [
  {
    id: 5,
    name: 'トマト',
    description: '3個',
    status: 'IN_CART',
    memberId: '1',
    categoryId: 3,
  },
  {
    id: 6,
    name: 'じゃがいも',
    description: '5個',
    status: 'IN_CART',
    memberId: '2',
    categoryId: 3,
  },
  {
    id: 7,
    name: '豚肉',
    description: '300g',
    status: 'IN_CART',
    memberId: '3',
    categoryId: 4,
  },
]

export interface InCartItemsTabProps {
  items?: ShoppingItem[]
  onMarkAsPurchased?: (item: ShoppingItem) => void
  onReturnToUnpurchased?: (item: ShoppingItem) => void
}

export const InCartItemsTab = ({
  items = dummyItems, // 本来はpropsから受け取る
  onMarkAsPurchased,
  onReturnToUnpurchased,
}: InCartItemsTabProps) => {
  // アイテムがない場合は空の表示
  if (items.length === 0) {
    return (
      <View style={styles.emptyContainer}>
        <Text style={styles.emptyText}>カート内のアイテムがありません</Text>
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
            <ActionButtonContainer>
              {/* 未購入に戻すボタン */}
              <ActionButton
                icon={
                  <CornerDownLeft
                    width={20}
                    height={20}
                    stroke={Colors.subText}
                  />
                }
                label="戻す"
                status="IN_CART"
                kind="return"
                onPress={() => onReturnToUnpurchased?.(item)}
              />

              {/* 購入済みにするボタン */}
              <ActionButton
                icon={<Check width={20} height={20} stroke={Colors.white} />}
                label="購入済"
                status="IN_CART"
                kind="move"
                onPress={() => onMarkAsPurchased?.(item)}
              />
            </ActionButtonContainer>
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
