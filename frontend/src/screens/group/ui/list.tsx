import { GroupList } from '@/features/group'
import { View, StyleSheet } from 'react-native'
import { NativeStackScreenProps } from '@react-navigation/native-stack'
import { GroupStackParamList } from './stack-navigator'

type Props = NativeStackScreenProps<GroupStackParamList, 'GroupList'>

export const GroupListScreen = ({ navigation }: Props) => {
  const handleGroupPress = (groupId: string, groupName: string) => {
    navigation.navigate('GroupDetail', { groupId, groupName })
  }

  return (
    <View style={styles.container}>
      <GroupList onPress={handleGroupPress} />
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    padding: 16,
  },
})
