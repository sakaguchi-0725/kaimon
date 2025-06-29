import { GroupList } from '@/features/group'
import { useJoinedGroups } from '@/features/group/lib/use-joined-groups'
import { View, StyleSheet, ActivityIndicator, Text } from 'react-native'
import { NativeStackScreenProps } from '@react-navigation/native-stack'
import { GroupStackParamList } from './stack-navigator'
import { Colors } from '@/shared/constants'

type Props = NativeStackScreenProps<GroupStackParamList, 'GroupList'>

export const GroupListScreen = ({ navigation }: Props) => {
  const { groups, isLoading, error } = useJoinedGroups()

  const handleGroupPress = (groupId: string, groupName: string) => {
    navigation.navigate('GroupDetail', { groupId, groupName })
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
})
