import { createNativeStackNavigator } from "@react-navigation/native-stack"
import { WelcomeScreen } from "./welcome";
import { SignUpScreen } from "./signup";
import { LoginScreen } from "./login";

export type AuthStackParamList = {
  Welcome: undefined;
  SignUp: undefined
  Login: undefined
};

const Stack = createNativeStackNavigator<AuthStackParamList>()

export const AuthStackNavigator = () => {
  return (
    <Stack.Navigator screenOptions={{ headerShown: false }}>
      <Stack.Screen name="Welcome" component={WelcomeScreen} />
      <Stack.Screen name="SignUp" component={SignUpScreen} />
      <Stack.Screen name="Login" component={LoginScreen} />
    </Stack.Navigator>
  )
}