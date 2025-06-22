import { View, Text, StyleSheet, TouchableOpacity } from 'react-native'
import { Colors } from '@/shared/constants'
import { createMaterialTopTabNavigator } from '@react-navigation/material-top-tabs'
import { Plus } from 'react-native-feather'
import { useModal } from '@/shared/ui/modal'
import { UnpurchasedItemsTab, InCartItemsTab, PurchasedItemsTab, useRealtimeShopping } from '@/features/shopping'
import { STATUS_LABELS, TAB_NAMES, TabInfo } from '@/features/shopping/model/constants'

const Tab = createMaterialTopTabNavigator()

export interface RealtimeShoppingListScreenProps {
  route?: {
    params?: {
      groupId?: string;
      groupName?: string;
    }
  }
}

export const RealtimeShoppingListScreen = ({ route }: RealtimeShoppingListScreenProps) => {
  // 選択されたグループの情報を取得
  const groupId = route?.params?.groupId
  const groupName = route?.params?.groupName || '買い物リスト'
  
  // リアルタイム買い物フック
  const { 
    isLoading,
    error,
    unpurchasedItems,
    inCartItems,
    purchasedItems,
    addToCart,
    markAsPurchased,
    returnToUnpurchased,
    returnToCart
  } = useRealtimeShopping({ groupId })
  
  // アイテム追加モーダル用
  const { isVisible: addItemModalVisible, openModal: openAddItemModal, closeModal: closeAddItemModal } = useModal()

  // タブ情報の定義
  const tabInfos: TabInfo[] = [
    {
      name: TAB_NAMES.UNPURCHASED,
      component: () => (
        <UnpurchasedItemsTab
          items={unpurchasedItems}
          onAddToCart={addToCart}
        />
      ),
      options: { tabBarLabel: STATUS_LABELS.UNPURCHASED }
    },
    {
      name: TAB_NAMES.IN_CART,
      component: () => (
        <InCartItemsTab
          items={inCartItems}
          onMarkAsPurchased={markAsPurchased}
          onReturnToUnpurchased={returnToUnpurchased}
        />
      ),
      options: { tabBarLabel: STATUS_LABELS.IN_CART }
    },
    {
      name: TAB_NAMES.PURCHASED,
      component: () => (
        <PurchasedItemsTab
          items={purchasedItems}
          onReturnToCart={returnToCart}
        />
      ),
      options: { tabBarLabel: STATUS_LABELS.PURCHASED }
    }
  ]

  // エラーがあれば表示
  if (error) {
    return (
      <View style={styles.errorContainer}>
        <Text style={styles.errorText}>{error}</Text>
      </View>
    )
  }

  return (
    <View style={styles.container}>
      <View style={styles.header}>
        <Text style={styles.title}>{groupName}</Text>
        <Text style={styles.subtitle}>
          {groupId ? `ID: ${groupId}` : 'グループが選択されていません'}
        </Text>
      </View>
      
      <Tab.Navigator
        screenOptions={{
          tabBarActiveTintColor: Colors.primary,
          tabBarInactiveTintColor: Colors.subText,
          tabBarIndicatorStyle: { backgroundColor: Colors.primary },
          tabBarLabelStyle: { fontSize: 14, fontWeight: 'bold' },
          tabBarStyle: { backgroundColor: Colors.white },
        }}
      >
        {tabInfos.map((tab) => (
          <Tab.Screen
            key={tab.name}
            name={tab.name}
            component={tab.component}
            options={tab.options}
          />
        ))}
      </Tab.Navigator>
      
      {/* 新規メモ追加ボタン */}
      <TouchableOpacity
        style={styles.addButton}
        onPress={openAddItemModal}
        activeOpacity={0.8}
      >
        <Plus width={24} height={24} stroke={Colors.white} />
      </TouchableOpacity>
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: Colors.backgroundGray,
  },
  header: {
    padding: 16,
    backgroundColor: Colors.white,
    borderBottomWidth: 1,
    borderBottomColor: Colors.border,
  },
  title: {
    fontSize: 20,
    fontWeight: 'bold',
    color: Colors.mainText,
  },
  subtitle: {
    fontSize: 14,
    color: Colors.subText,
    marginTop: 2,
  },
  errorContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: Colors.backgroundGray,
  },
  errorText: {
    color: Colors.error,
    fontSize: 16,
    textAlign: 'center',
    marginHorizontal: 20,
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
    elevation: 3,
  },
}) 