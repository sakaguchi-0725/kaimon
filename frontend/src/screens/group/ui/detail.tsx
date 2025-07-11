import React, { useState } from 'react'
import { StyleSheet, Text, View, TouchableOpacity } from 'react-native'
import { NativeStackScreenProps } from '@react-navigation/native-stack'
import { GroupStackParamList } from './stack-navigator'
import {
  GroupDetail,
  MemberList,
  useGroupDetail,
  GroupInvitationModal,
} from '@/features/group'
import { Container } from '@/shared/ui/container'
import { Colors } from '@/shared/constants'
import { UserPlus } from 'react-native-feather'

type Props = NativeStackScreenProps<GroupStackParamList, 'GroupDetail'>

export const GroupDetailScreen = ({ route }: Props) => {
  const { groupId } = route.params
  const { data: groupData, isLoading, error } = useGroupDetail(groupId)
  const [isInvitationModalVisible, setIsInvitationModalVisible] =
    useState(false)

  const handleAddMember = () => {
    setIsInvitationModalVisible(true)
  }

  const handleEdit = () => {
    console.log('編集ボタンがタップされました')
    // ここに編集の処理を実装
  }

  if (isLoading) {
    return (
      <Container>
        <Text>読み込み中...</Text>
      </Container>
    )
  }

  if (error || !groupData) {
    return (
      <Container>
        <Text>エラーが発生しました: {error}</Text>
      </Container>
    )
  }

  return (
    <>
      <Container>
        <GroupDetail data={groupData} onEdit={handleEdit} />
      </Container>

      <Container style={styles.membersContainer}>
        <View style={styles.sectionHeader}>
          <Text style={styles.sectionTitle}>メンバー</Text>
          <TouchableOpacity onPress={handleAddMember} style={styles.addButton}>
            <UserPlus stroke={Colors.primary} width={20} height={20} />
          </TouchableOpacity>
        </View>
        <MemberList members={groupData.members || []} />
      </Container>

      <GroupInvitationModal
        isVisible={isInvitationModalVisible}
        onClose={() => setIsInvitationModalVisible(false)}
      />
    </>
  )
}

const styles = StyleSheet.create({
  membersContainer: {
    flex: 1,
    backgroundColor: Colors.white,
    borderTopLeftRadius: 16,
    borderTopRightRadius: 16,
    paddingTop: 16,
  },
  sectionHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 12,
    paddingHorizontal: 4,
  },
  sectionTitle: {
    fontSize: 16,
    fontWeight: '600',
    color: Colors.mainText,
  },
  addButton: {
    padding: 4,
  },
})
