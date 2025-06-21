import { NavigationContainer } from "@react-navigation/native"
import { AppNavigator } from "./navigator"
import { DefaultTheme } from "@react-navigation/native"

const AppTheme = {
  ...DefaultTheme,
  colors: {
    ...DefaultTheme.colors,
    background: 'white',
  },
}

const App = () => {
  return (
    <NavigationContainer theme={AppTheme}>
      <AppNavigator />
    </NavigationContainer>
  )
}

export default App
