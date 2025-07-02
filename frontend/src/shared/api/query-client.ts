import createFetchClient from 'openapi-fetch'
import createClient from 'openapi-react-query'
import type { paths } from './schema'
import { getAuth } from '@react-native-firebase/auth'
import { isApiError, handleUnexpectedError } from './error'

const baseUrl: string = process.env.EXPO_PUBLIC_API_BASE_URL || ''

const fetchClient = createFetchClient<paths>({ baseUrl })

fetchClient.use({
  async onRequest({ request }) {
    const auth = getAuth()
    const user = auth.currentUser
    if (user) {
      const token = await user.getIdToken()
      request.headers.set('Authorization', `Bearer ${token}`)
    }

    console.log('API Request:', {
      url: request.url,
      method: request.method,
      headers: Object.fromEntries(request.headers.entries()),
    })

    return request
  },
  onResponse({ response }) {
    console.log('API Response:', {
      url: response.url,
      status: response.status,
      body: JSON.stringify(response.body),
    })
  },
  onError({ error }) {
    console.error('API Error:', error)
    if (!isApiError(error)) {
      handleUnexpectedError(error)
    }
  },
})

const $api = createClient(fetchClient)

const api = {
  useQuery: $api.useQuery,
  useMutation: $api.useMutation,
  useInfiniteQuery: $api.useInfiniteQuery,
  useSuspenseQuery: $api.useSuspenseQuery,
}

export default api
export { api }
