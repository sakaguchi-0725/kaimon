import React, { ReactNode } from 'react';
import { StyleSheet, Text, View, ViewStyle } from 'react-native';
import { Dropdown as RNDropdown } from 'react-native-element-dropdown';
import { Colors } from '@/shared/constants';
import { DropdownOption } from '../model';

type Props<T> = {
  selectedValue: T | null;
  onValueChange: (value: T) => void;
  options: DropdownOption<T>[];
  placeholder?: string;
  label?: string;
  required?: boolean;
  containerStyle?: ViewStyle;
  disabled?: boolean;
}

export function Dropdown<T>({
  selectedValue,
  onValueChange,
  options,
  placeholder = '選択してください',
  label,
  required = false,
  containerStyle,
  disabled = false,
}: Props<T>): ReactNode {
  return (
    <View>
      {label && (
        <Text style={styles.label}>
          {label} {required && <Text style={styles.required}>*</Text>}
        </Text>
      )}
      <RNDropdown
        style={[styles.dropdown, disabled && styles.disabledDropdown]}
        placeholderStyle={styles.placeholderStyle}
        selectedTextStyle={styles.selectedTextStyle}
        containerStyle={styles.dropdownContainer}
        data={options}
        labelField="label"
        valueField="value"
        placeholder={placeholder}
        value={selectedValue}
        onChange={(item) => {
          onValueChange(item.value as T);
        }}
        disable={disabled}
        maxHeight={150}
        renderItem={(item, selected) => (
          <View style={[styles.item, selected && styles.selectedItem]}>
            <Text style={selected ? styles.selectedItemText : styles.itemText}>
              {item.label}
            </Text>
          </View>
        )}
      />
    </View>
  );
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
  dropdown: {
    backgroundColor: Colors.backgroundGray,
    borderRadius: 8,
    padding: 12,
    height: 50,
  },
  dropdownContainer: {
    borderRadius: 8,
    borderWidth: 1,
    borderColor: Colors.border,
  },
  disabledDropdown: {
    opacity: 0.5,
  },
  placeholderStyle: {
    fontSize: 16,
    color: Colors.subText + '80',
  },
  selectedTextStyle: {
    fontSize: 16,
    color: Colors.mainText,
  },
  item: {
    padding: 12,
    flexDirection: 'row',
    alignItems: 'center',
  },
  selectedItem: {
    backgroundColor: Colors.backgroundGray,
  },
  itemText: {
    fontSize: 16,
    color: Colors.mainText,
  },
  selectedItemText: {
    fontSize: 16,
    color: Colors.primary,
    fontWeight: '500',
  },
}); 