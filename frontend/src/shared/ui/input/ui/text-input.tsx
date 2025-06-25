import React from 'react'
import {
  View,
  Text,
  TextInput as RNTextInput,
  StyleSheet,
  TextInputProps as RNTextInputProps,
  ViewStyle,
} from 'react-native'
import { Colors } from '@/shared/constants'

type Props = Omit<RNTextInputProps, 'style'> & {
  label?: string
  required?: boolean
  error?: string
  containerStyle?: ViewStyle
  inputStyle?: ViewStyle
}

export const TextInput: React.FC<Props> = ({
  label,
  required = false,
  error,
  containerStyle,
  inputStyle,
  placeholder,
  ...props
}) => {
  return (
    <View style={containerStyle}>
      {label && (
        <Text style={styles.label}>
          {label} {required && <Text style={styles.required}>*</Text>}
        </Text>
      )}
      <RNTextInput
        style={[styles.input, inputStyle]}
        placeholder={placeholder}
        placeholderTextColor={Colors.subText + '80'}
        {...props}
      />
      {error && <Text style={styles.error}>{error}</Text>}
    </View>
  )
}

const styles = StyleSheet.create({
  label: {
    fontSize: 14,
    fontWeight: '500',
    color: Colors.mainText,
    marginBottom: 8,
  },
  required: {
    color: Colors.error,
  },
  input: {
    backgroundColor: Colors.backgroundGray,
    borderRadius: 8,
    padding: 12,
    fontSize: 16,
    color: Colors.mainText,
  },
  error: {
    fontSize: 12,
    color: Colors.error,
    marginTop: 4,
  },
})
