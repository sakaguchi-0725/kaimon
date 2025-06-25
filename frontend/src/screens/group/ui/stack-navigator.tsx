import React from 'react'
import { createNativeStackNavigator } from '@react-navigation/native-stack'
import { GroupListScreen } from './list'
import { GroupDetailScreen } from './detail'
import { Colors } from '@/shared/constants'

export type GroupStackParamList = {
  GroupList: undefined
  GroupDetail: {
    groupId: string
    groupName: string
  }
}

const Stack = createNativeStackNavigator<GroupStackParamList>()

export const GroupStackNavigator = () => {
  return (
    <Stack.Navigator
      screenOptions={{
        headerBackTitle: '戻る',
        headerTintColor: Colors.primary,
      }}
    >
      <Stack.Screen
        name="GroupList"
        component={GroupListScreen}
        options={{
          title: 'グループ一覧',
          headerTitleStyle: {
            color: Colors.mainText,
          },
          contentStyle: {
            backgroundColor: Colors.backgroundGray,
          },
        }}
      />
      <Stack.Screen
        name="GroupDetail"
        component={GroupDetailScreen}
        options={({ route }) => ({
          title: route.params.groupName,
          headerTitleStyle: {
            color: Colors.mainText,
          },
          contentStyle: {
            backgroundColor: Colors.backgroundGray,
          },
        })}
      />
    </Stack.Navigator>
  )
}
