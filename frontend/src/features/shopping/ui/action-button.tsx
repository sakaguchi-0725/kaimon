import { ReactNode } from 'react'
import { View, Text, StyleSheet, TouchableOpacity } from 'react-native'
import { Colors } from '@/shared/constants'
import { ShoppingItemStatus } from '../model'

// アクションの種類を表す型
export type ActionKind = 'add' | 'move' | 'return' | 'default'

export interface ActionButtonProps {
  onPress?: () => void
  icon: ReactNode
  label: string
  status?: ShoppingItemStatus
  kind?: ActionKind
  style?: any
}

// ステータスに応じた色を取得
const getColorByStatus = (status: ShoppingItemStatus, kind: ActionKind): string => {
  // 戻る系のアクション
  if (kind === 'return') {
    return Colors.secondary
  }

  switch (status) {
    case 'UNPURCHASED':
      return kind === 'add' ? Colors.primary : Colors.subText
    case 'IN_CART':
      return kind === 'move' ? Colors.success : Colors.subText
    case 'PURCHASED':
      return Colors.secondary
    default:
      return Colors.primary
  }
}

// テキスト色を取得
const getTextColor = (kind: ActionKind): string => {
  return kind === 'return' ? Colors.subText : Colors.white
}

export const ActionButton = ({ 
  onPress, 
  icon, 
  label, 
  status = 'UNPURCHASED',
  kind = 'default',
  style
}: ActionButtonProps) => {
  // ステータスとアクションの種類に応じた色を取得
  const backgroundColor = getColorByStatus(status, kind)
  const textColor = getTextColor(kind)

  return (
    <TouchableOpacity
      style={[
        styles.actionButton,
        { backgroundColor },
        style
      ]}
      onPress={onPress}
    >
      {icon}
      <Text style={[styles.actionButtonText, { color: textColor }]}>{label}</Text>
    </TouchableOpacity>
  )
}

export const ActionButtonContainer = ({ children }: { children: ReactNode }) => {
  return (
    <View style={styles.actionContainer}>
      {children}
    </View>
  )
}

const styles = StyleSheet.create({
  actionContainer: {
    flexDirection: 'row',
    marginLeft: 8,
  },
  actionButton: {
    borderRadius: 4,
    paddingHorizontal: 12,
    paddingVertical: 8,
    flexDirection: 'row',
    alignItems: 'center',
    marginLeft: 8,
  },
  actionButtonText: {
    fontSize: 14,
    fontWeight: '500',
    marginLeft: 4,
  },
}) 