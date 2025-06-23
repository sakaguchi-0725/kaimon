import React from 'react'
import { StyleSheet } from 'react-native'
import { TextInput } from './text-input'

type Props = {
  label?: string
  required?: boolean
  error?: string
  containerStyle?: any
  inputStyle?: any
  value?: string
  onChangeText?: (text: string) => void
  placeholder?: string
  numberOfLines?: number
}

export const Textarea: React.FC<Props> = ({
  numberOfLines = 3,
  inputStyle,
  ...props
}) => {
  return (
    <TextInput
      multiline
      numberOfLines={numberOfLines}
      inputStyle={{
        ...StyleSheet.flatten(styles.textarea),
        ...StyleSheet.flatten(inputStyle || {})
      }}
      {...props}
    />
  )
}

const styles = StyleSheet.create({
  textarea: {
    height: 80,
    textAlignVertical: 'top',
  },
}) 