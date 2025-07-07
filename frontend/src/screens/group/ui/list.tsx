import { GroupList, useJoinedGroups, CreateGroupModal } from '@/features/group'
import {
  View,
  StyleSheet,
  ActivityIndicator,
  Text,
  TouchableOpacity,
} from 'react-native'
import { NativeStackScreenProps } from '@react-navigation/native-stack'
import { GroupStackParamList } from './stack-navigator'
import { Colors } from '@/shared/constants'
import { useModal, BottomSheetModal } from '@/shared/ui/modal'
import { Plus } from 'react-native-feather'

type Props = NativeStackScreenProps<GroupStackParamList, 'GroupList'>

export const GroupListScreen = ({ navigation }: Props) => {
  const { groups, isLoading, error, refetch } = useJoinedGroups()
  const {
    isVisible: actionModalVisible,
    openModal: openActionModal,
    closeModal: closeActionModal,
  } = useModal()
  const {
    isVisible: createModalVisible,
    openModal: openCreateModal,
    closeModal: closeCreateModal,
  } = useModal()

  const handleGroupPress = (groupId: string, groupName: string) => {
    navigation.navigate('GroupDetail', { groupId, groupName })
  }

  const handleCreateGroup = () => {
    closeActionModal()
    openCreateModal()
  }

  const handleJoinGroup = () => {
    closeActionModal()
    // TODO: 本番環境では適切なログ管理ツールを使用
    console.log('グループに参加')
  }

  const handleCreateSuccess = () => {
    // TODO: 本番環境では適切なログ管理ツールを使用
    console.log('グループ作成成功')
  }

  if (isLoading) {
    return (
      <View style={styles.centerContainer}>
        <ActivityIndicator size="large" color={Colors.primary} />
        <Text style={styles.loadingText}>読み込み中...</Text>
      </View>
    )
  }

  if (error) {
    return (
      <View style={styles.centerContainer}>
        <Text style={styles.errorText}>{error}</Text>
      </View>
    )
  }

  return (
    <View style={styles.container}>
      <GroupList groups={groups} onPress={handleGroupPress} />

      {/* 新規グループ追加ボタン */}
      <TouchableOpacity
        style={styles.addButton}
        onPress={openActionModal}
        activeOpacity={0.8}
        accessible={true}
        accessibilityLabel="グループを追加"
        accessibilityHint="新しいグループを作成するか、既存のグループに参加できます"
      >
        <Plus width={24} height={24} stroke={Colors.white} />
      </TouchableOpacity>

      {/* アクション選択モーダル */}
      <BottomSheetModal
        isVisible={actionModalVisible}
        onClose={closeActionModal}
      >
        <View style={styles.actionModalContent}>
          <TouchableOpacity
            style={styles.actionButton}
            onPress={handleCreateGroup}
            activeOpacity={0.8}
            accessible={true}
            accessibilityLabel="グループを作成"
            accessibilityRole="button"
          >
            <Text style={styles.actionButtonText}>グループを作成</Text>
          </TouchableOpacity>
          <TouchableOpacity
            style={styles.actionButton}
            onPress={handleJoinGroup}
            activeOpacity={0.8}
            accessible={true}
            accessibilityLabel="グループに参加"
            accessibilityRole="button"
          >
            <Text style={styles.actionButtonText}>グループに参加</Text>
          </TouchableOpacity>
        </View>
      </BottomSheetModal>

      {/* グループ作成モーダル */}
      <CreateGroupModal
        isVisible={createModalVisible}
        onClose={closeCreateModal}
        onSuccess={handleCreateSuccess}
        onRefetch={refetch}
      />
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 16,
  },
  centerContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    padding: 20,
  },
  loadingText: {
    marginTop: 12,
    fontSize: 16,
    color: Colors.subText,
  },
  errorText: {
    fontSize: 16,
    color: Colors.error,
    textAlign: 'center',
  },
  addButton: {
    position: 'absolute',
    bottom: 24,
    right: 24,
    width: 56,
    height: 56,
    borderRadius: 28,
    backgroundColor: Colors.primary,
    justifyContent: 'center',
    alignItems: 'center',
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 1 },
    shadowOpacity: 0.2,
    shadowRadius: 2,
  },
  actionModalContent: {
    paddingVertical: 10,
  },
  actionButton: {
    paddingVertical: 16,
    paddingHorizontal: 20,
    borderRadius: 8,
    backgroundColor: Colors.white,
    marginVertical: 4,
  },
  actionButtonText: {
    fontSize: 16,
    color: Colors.mainText,
    textAlign: 'center',
    fontWeight: '500',
  },
})
