import React from 'react';
import { StyleSheet, Text, TouchableOpacity, View, StyleProp, ViewStyle, TextStyle } from 'react-native';
import { Colors } from '@/shared/constants';
import { ButtonSize, ButtonVariant, ButtonColor } from '../model';

type Props = {
  text: string;
  onPress: () => void;
  size?: ButtonSize;
  variant?: ButtonVariant;
  color?: ButtonColor;
  icon?: React.ReactNode;
  iconPosition?: 'left' | 'right';
  disabled?: boolean;
  style?: StyleProp<ViewStyle>;
  textStyle?: StyleProp<TextStyle>;
}

export const Button: React.FC<Props> = ({
  text,
  onPress,
  size = 'full',
  variant = 'solid',
  color = 'primary',
  icon,
  iconPosition = 'left',
  disabled = false,
  style,
  textStyle,
}) => {
  const getBackgroundColor = () => {
    if (disabled) return Colors.secondary;
    if (variant === 'solid') {
      return color === 'primary' ? Colors.primary : Colors.secondary;
    }
    return 'transparent';
  };

  const getBorderColor = () => {
    if (disabled) return Colors.secondary;
    if (variant === 'outline') {
      return color === 'primary' ? Colors.primary : Colors.secondary;
    }
    return 'transparent';
  };

  const getTextColor = () => {
    if (disabled) return Colors.subText;
    if (variant === 'solid') {
      return Colors.white;
    }
    return color === 'primary' ? Colors.primary : Colors.subText;
  };

  const buttonStyles = [
    styles.button,
    size === 'sm' ? styles.buttonSm : styles.buttonFull,
    variant === 'outline' && styles.buttonOutline,
    {
      backgroundColor: getBackgroundColor(),
      borderColor: getBorderColor(),
    },
    disabled && styles.buttonDisabled,
    style,
  ];

  const textStyles = [
    styles.text,
    size === 'sm' && styles.textSm,
    { color: getTextColor() },
    textStyle,
  ];

  return (
    <TouchableOpacity
      style={buttonStyles}
      onPress={onPress}
      disabled={disabled}
      activeOpacity={0.7}
    >
      <View style={styles.contentContainer}>
        {icon && iconPosition === 'left' && <View style={styles.iconLeft}>{icon}</View>}
        <Text style={textStyles}>{text}</Text>
        {icon && iconPosition === 'right' && <View style={styles.iconRight}>{icon}</View>}
      </View>
    </TouchableOpacity>
  );
};

const styles = StyleSheet.create({
  button: {
    borderRadius: 8,
    paddingVertical: 12,
    alignItems: 'center',
    justifyContent: 'center',
    borderWidth: 1,
    borderColor: 'transparent',
  },
  buttonSm: {
    paddingHorizontal: 16,
    paddingVertical: 8,
    alignSelf: 'flex-start',
  },
  buttonFull: {
    width: '100%',
  },
  buttonOutline: {
    backgroundColor: 'transparent',
  },
  buttonDisabled: {
    opacity: 0.5,
  },
  contentContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
  },
  text: {
    fontSize: 16,
    fontWeight: 'bold',
    textAlign: 'center',
  },
  textSm: {
    fontSize: 14,
  },
  iconLeft: {
    marginRight: 8,
  },
  iconRight: {
    marginLeft: 8,
  },
});
