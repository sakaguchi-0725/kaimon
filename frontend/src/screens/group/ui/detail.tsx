import React from 'react'
import { StyleSheet, Text, View, TouchableOpacity } from "react-native"
import { NativeStackScreenProps } from "@react-navigation/native-stack"
import { GroupStackParamList } from "./stack-navigator"
import { GroupDetail } from "@/features/group"
import { MemberList } from "@/features/member"
import { Container } from '@/shared/ui/container'
import { Colors } from '@/shared/constants'
import { UserPlus } from 'react-native-feather'

type Props = NativeStackScreenProps<GroupStackParamList, 'GroupDetail'>

export const GroupDetailScreen = ({ route }: Props) => {
  const { groupId, groupName } = route.params
  
  const handleAddMember = () => {
    console.log('メンバー追加ボタンがタップされました')
    // ここにメンバー追加の処理を実装
  }
  
  return (
    <>
      <Container>
        <GroupDetail />
      </Container>
      
      <Container style={styles.membersContainer}>
        <View style={styles.sectionHeader}>
          <Text style={styles.sectionTitle}>メンバー</Text>
          <TouchableOpacity onPress={handleAddMember} style={styles.addButton}>
            <UserPlus stroke={Colors.primary} width={20} height={20} />
          </TouchableOpacity>
        </View>
        <MemberList />
      </Container>
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
  }
})