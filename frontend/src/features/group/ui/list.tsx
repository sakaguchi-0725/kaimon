import { FlatList, StyleSheet, Text, TouchableOpacity, View } from "react-native"
import { JoinedGroup } from "../model"
import { ChevronRight } from "react-native-feather"
import { Colors } from "@/shared/constants"

const groups: JoinedGroup[] = [
  { id: '1', name: 'グループ1', memberCount: 2 },
  { id: '2', name: 'グループ2', memberCount: 5 },
  { id: '3', name: 'グループ3', memberCount: 3 },
]

type Props = {
  onPress?: (groupId: string, groupName: string) => void;
}

export const GroupList = ({ onPress }: Props) => {
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
})