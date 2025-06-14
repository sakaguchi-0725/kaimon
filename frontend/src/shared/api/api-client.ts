import createClient, { type Middleware } from 'openapi-fetch'
import type { paths } from './schema'
import type { ApiSchema } from './response'
import { API_URL } from '../constants'
import { showError } from '../ui/toast'

// エラーレスポンスの型定義
type ErrorResponse = ApiSchema<'Error'>

// エラーレスポンスかどうかを判定する型ガード関数
const isErrorResponse = (obj: unknown): obj is ErrorResponse => {
  return typeof obj === 'object' && obj !== null && 'code' in obj && 'message' in obj
}

const errorMiddleware: Middleware = {
  onError: async ({ error }) => {
    let errorObj: unknown = error
    let statusCode: number | undefined = undefined

    // エラーがResponseの場合
    if (error instanceof Response) {
      statusCode = error.status

      try {
        errorObj = await error.clone().json()
      } catch (e) {
        console.error('Failed to parse error response:', e)
      }
    }

    if (statusCode === 401) {
      showError('認証エラー', '再度ログインしてください')
      window.location.href = '/login'
      return
    }

    if (statusCode === 500) {
      showError('サーバーエラー', 'サーバーでエラーが発生しました')
      window.location.href = '/error/500'
      return
    }

    if (isErrorResponse(errorObj)) {
      showError('エラーが発生しました', errorObj.message)
      return
    }

    showError('予期せぬエラー', 'システムエラーが発生しました')
    window.location.href = '/error/500'
  },
}

const instance = createClient<paths>({
  baseUrl: API_URL,
  credentials: 'include',
})

instance.use(errorMiddleware)

export const client = {
  get: instance.GET,
  post: instance.POST,
  put: instance.PUT,
  delete: instance.DELETE,
}
