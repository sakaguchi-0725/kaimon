import { ReactNode } from 'react'
import { StyleSheet, TouchableOpacity, View, ViewStyle } from 'react-native'
import { Colors } from '@/shared/constants'

export interface CardProps {
  children: ReactNode
  onPress?: () => void
  style?: ViewStyle
  contentStyle?: ViewStyle
  activeOpacity?: number
}

export const Card = ({ 
  children, 
  onPress, 
  style, 
  contentStyle,
  activeOpacity = 0.7 
}: CardProps) => {
  // タッチ可能なカードかどうかでコンポーネントを分岐
  if (onPress) {
    return (
      <TouchableOpacity
        style={[styles.container, style]}
        onPress={onPress}
        activeOpacity={activeOpacity}
      >
        <View style={[styles.content, contentStyle]}>
          {children}
        </View>
      </TouchableOpacity>
    )
  }

  // タッチ不可のカード
  return (
    <View style={[styles.container, style]}>
      <View style={[styles.content, contentStyle]}>
        {children}
      </View>
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    backgroundColor: Colors.white,
    borderRadius: 8,
    borderColor: Colors.border,
    borderWidth: 0.5,
    marginHorizontal: 16,
    marginVertical: 4,
  },
  content: {
    padding: 16,
  },
}) 