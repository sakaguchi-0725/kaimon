import { NavigationContainer, DefaultTheme } from '@react-navigation/native'
import { AppNavigator } from './navigator'
import { GestureHandlerRootView } from 'react-native-gesture-handler'
import { StyleSheet } from 'react-native'
import { QueryClientProvider } from '@tanstack/react-query'
import { queryClient } from '@/shared/api'
import { AccountContextProvider } from './context'

const AppTheme = {
  ...DefaultTheme,
  colors: {
    ...DefaultTheme.colors,
    background: 'white',
  },
}

const App = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <AccountContextProvider>
        <GestureHandlerRootView style={styles.container}>
          <NavigationContainer theme={AppTheme}>
            <AppNavigator />
          </NavigationContainer>
        </GestureHandlerRootView>
      </AccountContextProvider>
    </QueryClientProvider>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
})

export default App
