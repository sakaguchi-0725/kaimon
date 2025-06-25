import React from 'react'
import { createNativeStackNavigator } from '@react-navigation/native-stack'
import { ShoppingItemListScreen } from './list'
import { RealtimeShoppingListScreen } from './realtime-list'
import { Colors } from '@/shared/constants'
import { TouchableOpacity, Text, Alert } from 'react-native'

export type ShoppingStackParamList = {
  ShoppingList: undefined
  RealtimeShopping: {
    groupId?: string
  }
}

const Stack = createNativeStackNavigator<ShoppingStackParamList>()

export const ShoppingStackNavigator = () => {
  return (
    <Stack.Navigator
      screenOptions={{
        headerBackTitle: '戻る',
        headerTintColor: Colors.primary,
      }}
    >
      <Stack.Screen
        name="ShoppingList"
        component={ShoppingItemListScreen}
        options={({ navigation }) => ({
          title: '買い物一覧',
          headerTitleStyle: {
            color: Colors.mainText,
          },
          contentStyle: {
            backgroundColor: Colors.backgroundGray,
          },
          // 右側に買い物開始ボタンを表示
          headerRight: () => (
            <TouchableOpacity
              onPress={() => {
                // 実際の実装では選択されているグループIDを取得
                const selectedGroupId = 'dummy-group-id' // 仮のID
                navigation.navigate('RealtimeShopping', {
                  groupId: selectedGroupId,
                })
              }}
              style={{ marginRight: 8 }}
            >
              <Text style={{ color: Colors.primary, fontWeight: 'bold' }}>
                買い物を始める
              </Text>
            </TouchableOpacity>
          ),
        })}
      />
      <Stack.Screen
        name="RealtimeShopping"
        component={RealtimeShoppingListScreen}
        options={({ navigation }) => ({
          title: 'リアルタイム買い物',
          headerTitleStyle: {
            color: Colors.mainText,
          },
          contentStyle: {
            backgroundColor: Colors.backgroundGray,
          },
          // 左側に終了ボタンを表示
          headerLeft: () => (
            <TouchableOpacity
              onPress={() => {
                Alert.alert(
                  '買い物を終了しますか？',
                  '現在の状態が保存されます。',
                  [
                    { text: 'キャンセル', style: 'cancel' },
                    {
                      text: '終了する',
                      onPress: () => navigation.goBack(),
                    },
                  ],
                )
              }}
              style={{ marginRight: 8 }}
            >
              <Text style={{ color: Colors.primary }}>終了</Text>
            </TouchableOpacity>
          ),
        })}
      />
    </Stack.Navigator>
  )
}
