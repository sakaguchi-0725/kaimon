import { FlatList, StyleSheet, Text, View } from 'react-native'
import { Member } from '../model'
import { MemberRoleIcon } from './role-icon'
import { StatusLabel } from './status-label'
import { Colors } from '@/shared/constants'

const members: Member[] = [
  { id: '1', name: 'メンバー1', role: 'admin', status: 'active' },
  { id: '2', name: 'メンバー2', role: 'member', status: 'active' },
  { id: '3', name: 'メンバー3', role: 'member', status: 'pending' },
]

export const MemberList = () => {
  const renderItem = (member: Member) => {
    return (
      <View style={styles.item}>
        <View style={styles.nameArea}>
          <MemberRoleIcon role={member.role} />
          <Text style={styles.nameText}>{member.name}</Text>
        </View>
        <StatusLabel status={member.status} />
      </View>
    )
  }

  return (
    <FlatList
      data={members}
      keyExtractor={(member) => member.id}
      renderItem={({ item }) => renderItem(item)}
    />
  )
}

const styles = StyleSheet.create({
  item: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    paddingVertical: 14,
    borderBottomWidth: 0.5,
    borderBottomColor: Colors.border,
  },
  nameArea: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: 10,
  },
  nameText: {
    fontSize: 16,
    fontWeight: '500',
    color: Colors.mainText,
  },
})
