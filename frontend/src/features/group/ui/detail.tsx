import React from 'react'
import { StyleSheet, Text, View, TouchableOpacity } from 'react-native'
import { GroupInfo } from '../model'
import { Colors } from '@/shared/constants'
import { Edit2 } from 'react-native-feather'

export const GroupDetail = () => {
  const data: GroupInfo = {
    id: '1',
    name: 'グループ1',
    description: 'テスト用のグループです',
    createdAt: '2025-01-02',
  }

  const handleEdit = () => {
    // 編集処理
    console.log('編集ボタンがタップされました')
  }

  return (
    <View style={styles.container}>
      <View style={styles.card}>
        <View style={styles.section}>
          <View style={styles.headerRow}>
            <Text style={styles.label}>グループ概要</Text>
            <TouchableOpacity onPress={handleEdit} style={styles.editButton}>
              <Edit2 stroke={Colors.primary} width={16} height={16} />
            </TouchableOpacity>
          </View>
          <Text style={styles.description}>{data.description}</Text>
        </View>

        <View style={styles.divider} />

        <View style={styles.section}>
          <View style={styles.infoRow}>
            <Text style={styles.infoLabel}>作成日</Text>
            <Text style={styles.infoValue}>{data.createdAt}</Text>
          </View>
        </View>
      </View>
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    paddingVertical: 2,
  },
  card: {
    backgroundColor: Colors.white,
    borderRadius: 12,
    borderColor: Colors.border,
    borderWidth: 0.5,
    padding: 16,
  },
  section: {
    paddingVertical: 2,
  },
  headerRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 8,
  },
  label: {
    fontSize: 15,
    color: Colors.mainText,
    fontWeight: '600',
  },
  editButton: {
    padding: 4,
  },
  description: {
    fontSize: 16,
    color: Colors.mainText,
    lineHeight: 22,
  },
  divider: {
    height: 1,
    backgroundColor: Colors.border,
    marginVertical: 12,
  },
  infoRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    paddingVertical: 6,
  },
  infoLabel: {
    fontSize: 14,
    color: Colors.subText,
  },
  infoValue: {
    fontSize: 14,
    color: Colors.mainText,
    fontWeight: '500',
  },
})
