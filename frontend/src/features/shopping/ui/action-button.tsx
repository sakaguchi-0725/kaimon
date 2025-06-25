import { ReactNode } from 'react'
import { View, StyleSheet } from 'react-native'
import { Button } from '@/shared/ui'
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
const getButtonColor = (status: ShoppingItemStatus, kind: ActionKind): 'primary' | 'secondary' | 'success' => {
  // 戻る系のアクション
  if (kind === 'return') {
    return 'secondary'
  }

  switch (status) {
    case 'UNPURCHASED':
      return kind === 'add' ? 'primary' : 'secondary'
    case 'IN_CART':
      return kind === 'move' ? 'success' : 'secondary'
    case 'PURCHASED':
      return 'secondary'
    default:
      return 'primary'
  }
}

// ボタンのバリアントを取得
const getButtonVariant = (kind: ActionKind): 'solid' | 'outline' | 'text' => {
  return kind === 'return' ? 'outline' : 'solid'
}

export const ActionButton = ({ 
  onPress, 
  icon, 
  label, 
  status = 'UNPURCHASED',
  kind = 'default',
  style
}: ActionButtonProps) => {
  // ステータスとアクションの種類に応じた設定を取得
  const color = getButtonColor(status, kind)
  const variant = getButtonVariant(kind)
  
  // onPressがundefinedの場合は空の関数を渡す
  const handlePress = onPress || (() => {})
  
  return (
    <Button
      text={label}
      onPress={handlePress}
      icon={icon}
      iconPosition="left"
      size="sm"
      variant={variant}
      color={color}
      style={[styles.actionButton, style]}
      textStyle={styles.actionButtonText}
    />
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
    marginLeft: 8,
    paddingHorizontal: 12,
    paddingVertical: 8,
  },
  actionButtonText: {
    fontSize: 14,
    fontWeight: '500',
    marginLeft: 4,
  },
}) 