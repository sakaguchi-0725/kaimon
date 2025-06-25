import {
  FlatList,
  StyleSheet,
  Text,
  TouchableOpacity,
  View,
} from 'react-native'
import { ShoppingItem, ShoppingItemStatus, STATUS_LABELS } from '../model'
import { Colors } from '@/shared/constants'
import { Check, ShoppingCart } from 'react-native-feather'
import { Card } from '@/shared/ui'

export interface ShoppingItemListProps {
  items: ShoppingItem[]
  onItemPress?: (item: ShoppingItem) => void
  onStatusChange?: (item: ShoppingItem, newStatus: ShoppingItemStatus) => void
}

export const ShoppingItemList = ({
  items,
  onItemPress,
  onStatusChange,
}: ShoppingItemListProps) => {
  // アイテムがない場合は空の表示
  if (items.length === 0) {
    return (
      <View style={styles.emptyContainer}>
        <Text style={styles.emptyText}>買い物リストがありません</Text>
      </View>
    )
  }

  // ステータスに応じたアイコンを表示
  const renderStatusIcon = (status: ShoppingItemStatus) => {
    switch (status) {
      case 'PURCHASED':
        return <Check width={20} height={20} stroke={Colors.success} />
      case 'IN_CART':
        return <ShoppingCart width={20} height={20} stroke={Colors.primary} />
      default:
        return null
    }
  }

  // ステータス切り替えボタン
  const handleStatusChange = (item: ShoppingItem) => {
    if (!onStatusChange) return

    // ステータスを順番に切り替える
    const statuses: ShoppingItemStatus[] = [
      'UNPURCHASED',
      'IN_CART',
      'PURCHASED',
    ]
    const currentIndex = statuses.indexOf(item.status)
    const nextIndex = (currentIndex + 1) % statuses.length
    onStatusChange(item, statuses[nextIndex])
  }

  return (
    <FlatList
      data={items}
      keyExtractor={(item) => item.id.toString()}
      renderItem={({ item }) => (
        <Card
          onPress={() => onItemPress?.(item)}
          contentStyle={styles.cardContent}
        >
          {/* ステータスアイコン */}
          <TouchableOpacity
            style={styles.statusButton}
            onPress={() => handleStatusChange(item)}
          >
            {renderStatusIcon(item.status)}
          </TouchableOpacity>

          {/* アイテム情報 */}
          <View style={styles.itemContent}>
            <Text style={styles.itemName}>{item.name}</Text>
            {item.description ? (
              <Text style={styles.itemDescription} numberOfLines={2}>
                {item.description}
              </Text>
            ) : null}
          </View>

          {/* ステータスラベル */}
          <View style={[styles.statusBadge, getStatusStyle(item.status)]}>
            <Text style={styles.statusText}>{STATUS_LABELS[item.status]}</Text>
          </View>
        </Card>
      )}
      contentContainerStyle={styles.listContent}
      showsVerticalScrollIndicator={false}
    />
  )
}

// ステータスに応じたスタイルを取得
const getStatusStyle = (status: ShoppingItemStatus) => {
  switch (status) {
    case 'PURCHASED':
      return styles.purchasedStatus
    case 'IN_CART':
      return styles.inCartStatus
    default:
      return styles.unpurchasedStatus
  }
}

const styles = StyleSheet.create({
  emptyContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  emptyText: {
    fontSize: 16,
    color: Colors.subText,
  },
  listContent: {
    paddingVertical: 8,
  },
  cardContent: {
    flexDirection: 'row',
    alignItems: 'center',
    padding: 16,
  },
  statusButton: {
    width: 40,
    height: 40,
    borderRadius: 20,
    backgroundColor: Colors.backgroundGray,
    justifyContent: 'center',
    alignItems: 'center',
    marginRight: 12,
  },
  itemContent: {
    flex: 1,
  },
  itemName: {
    fontSize: 16,
    fontWeight: '500',
    color: Colors.mainText,
    marginBottom: 4,
  },
  itemDescription: {
    fontSize: 14,
    color: Colors.subText,
  },
  statusBadge: {
    paddingHorizontal: 8,
    paddingVertical: 4,
    borderRadius: 4,
    marginLeft: 8,
  },
  unpurchasedStatus: {
    backgroundColor: Colors.backgroundGray,
  },
  inCartStatus: {
    backgroundColor: Colors.primary + '20', // 透明度を下げた色
  },
  purchasedStatus: {
    backgroundColor: Colors.success + '20', // 透明度を下げた色
  },
  statusText: {
    fontSize: 12,
    fontWeight: '500',
  },
})
