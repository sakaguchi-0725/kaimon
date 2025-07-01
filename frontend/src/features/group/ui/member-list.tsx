import { FlatList, StyleSheet, Text, View } from 'react-native'
import { Member } from '../model'
import { MemberRoleIcon } from './member-role-icon'
import { StatusLabel } from './member-status-label'
import { Colors } from '@/shared/constants'

interface MemberListProps {
  members: Member[]
}

export const MemberList = ({ members }: MemberListProps) => {
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
      keyExtractor={(member) => member.id || ''}
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
