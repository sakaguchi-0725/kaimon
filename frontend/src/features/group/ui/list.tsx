import {
  FlatList,
  StyleSheet,
  Text,
  TouchableOpacity,
  View,
} from 'react-native'
import { JoinedGroup } from '../model'
import { ChevronRight } from 'react-native-feather'
import { Colors } from '@/shared/constants'

type Props = {
  groups: JoinedGroup[]
  onPress?: (groupId: string, groupName: string) => void
}

export const GroupList = ({ groups, onPress }: Props) => {
  const renderItem = (item: JoinedGroup) => (
    <TouchableOpacity
      style={styles.item}
      onPress={() => onPress?.(item.id, item.name)}
    >
      <View style={styles.titleArea}>
        <Text style={styles.title}>{item.name}</Text>
        <Text style={styles.description}>メンバー数: {item.memberCount}名</Text>
      </View>
      <ChevronRight stroke={Colors.primary} />
    </TouchableOpacity>
  )

  if (groups.length === 0) {
    return (
      <View style={styles.emptyContainer}>
        <Text style={styles.emptyTitle}>グループがまだありません</Text>
        <Text style={styles.emptyDescription}>
          グループを作成するか、グループに参加して{'\n'}
          みんなで買い物リストを共有しましょう！
        </Text>
      </View>
    )
  }

  return (
    <FlatList
      data={groups}
      keyExtractor={(item) => item.id}
      renderItem={({ item }) => renderItem(item)}
    />
  )
}

const styles = StyleSheet.create({
  item: {
    flex: 1,
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: 16,
    borderRadius: 8,
    backgroundColor: Colors.white,
    marginBottom: 12,
    borderWidth: 1,
    borderColor: '#D4D5D2',
  },
  titleArea: {
    gap: 2,
  },
  title: {
    fontSize: 15,
    fontWeight: 'bold',
    color: Colors.mainText,
  },
  description: {
    fontSize: 13,
    color: Colors.subText,
  },
  emptyContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    paddingHorizontal: 32,
    paddingVertical: 64,
  },
  emptyTitle: {
    fontSize: 18,
    fontWeight: 'bold',
    color: Colors.mainText,
    marginBottom: 12,
    textAlign: 'center',
  },
  emptyDescription: {
    fontSize: 14,
    color: Colors.subText,
    textAlign: 'center',
    lineHeight: 20,
  },
})
