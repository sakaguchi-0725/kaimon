import * as SecureStore from 'expo-secure-store'

const ACCESS_TOKEN_KEY = 'accessToken'
const REFRESH_TOKEN_KEY = 'refreshToken'

export const authStorage = {
  // AccessToken関連
  async setAccessToken(token: string): Promise<void> {
    await SecureStore.setItemAsync(ACCESS_TOKEN_KEY, token)
  },

  async getAccessToken(): Promise<string | null> {
    return await SecureStore.getItemAsync(ACCESS_TOKEN_KEY)
  },

  async removeAccessToken(): Promise<void> {
    await SecureStore.deleteItemAsync(ACCESS_TOKEN_KEY)
  },

  // RefreshToken関連
  async setRefreshToken(token: string): Promise<void> {
    await SecureStore.setItemAsync(REFRESH_TOKEN_KEY, token)
  },

  async getRefreshToken(): Promise<string | null> {
    return await SecureStore.getItemAsync(REFRESH_TOKEN_KEY)
  },

  async removeRefreshToken(): Promise<void> {
    await SecureStore.deleteItemAsync(REFRESH_TOKEN_KEY)
  },

  // 両方のトークンを一括削除
  async clearTokens(): Promise<void> {
    await Promise.all([
      SecureStore.deleteItemAsync(ACCESS_TOKEN_KEY),
      SecureStore.deleteItemAsync(REFRESH_TOKEN_KEY),
    ])
  },

  // 両方のトークンを一括保存
  async setTokens(accessToken: string, refreshToken: string): Promise<void> {
    await Promise.all([
      SecureStore.setItemAsync(ACCESS_TOKEN_KEY, accessToken),
      SecureStore.setItemAsync(REFRESH_TOKEN_KEY, refreshToken),
    ])
  },
}
