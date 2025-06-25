import { StyleSheet, View, ViewStyle } from 'react-native'
import React, { ReactNode } from 'react'
import { Colors } from '@/shared/constants'

type ContainerProps = {
  children: ReactNode
  backgroundColor?: 'gray'
  style?: ViewStyle
}

export const Container: React.FC<ContainerProps> = ({
  children,
  backgroundColor,
  style,
}) => {
  const containerStyle: ViewStyle = {
    ...styles.container,
    backgroundColor: backgroundColor === 'gray' ? Colors.secondary : undefined,
    ...style,
  }

  return <View style={containerStyle}>{children}</View>
}

const styles = StyleSheet.create({
  container: {
    padding: 16,
  },
})
