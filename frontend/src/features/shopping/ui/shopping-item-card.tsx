import { ReactNode } from 'react'
import { View, Text, StyleSheet } from 'react-native'
import { Colors } from '@/shared/constants'
import { ShoppingItem } from '../model'
import { Card } from '@/shared/ui'

export interface ShoppingItemCardProps {
  item: ShoppingItem
  actionButtons: ReactNode
}

export const ShoppingItemCard = ({ item, actionButtons }: ShoppingItemCardProps) => {
  return (
    <Card contentStyle={styles.cardContent}>
      {/* アイテム情報 */}
      <View style={styles.itemContent}>
        <Text style={styles.itemName}>{item.name}</Text>
        {item.description ? (
          <Text style={styles.itemDescription} numberOfLines={2}>
            {item.description}
          </Text>
        ) : null}
      </View>

      {/* アクションボタン */}
      {actionButtons}
    </Card>
  )
}

const styles = StyleSheet.create({
  cardContent: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    padding: 16,
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
}) 