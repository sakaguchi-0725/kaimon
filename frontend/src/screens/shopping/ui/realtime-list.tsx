import { View, Text, StyleSheet } from 'react-native'
import { Colors } from '@/shared/constants'

export interface RealtimeShoppingListScreenProps {
  route?: {
    params?: {
      groupId?: string
    }
  }
}

export const RealtimeShoppingListScreen = ({ route }: RealtimeShoppingListScreenProps) => {
  // 選択されたグループIDを取得
  const groupId = route?.params?.groupId

  return (
    <View style={styles.container}>
      <Text style={styles.title}>リアルタイム買い物</Text>
      <Text style={styles.subtitle}>
        {groupId ? `グループID: ${groupId}` : 'グループが選択されていません'}
      </Text>
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: Colors.backgroundGray,
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    color: Colors.mainText,
    marginBottom: 16,
  },
  subtitle: {
    fontSize: 16,
    color: Colors.subText,
  },
}) 