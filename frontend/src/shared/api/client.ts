import createClient from 'openapi-fetch'
import type { paths } from './schema'
import { getAuth } from '@react-native-firebase/auth'
import { isApiError, handleUnexpectedError } from './error'

let baseUrl: string = process.env.EXPO_PUBLIC_API_BASE_URL || ''

const instance = createClient<paths>({ baseUrl })

instance.use({
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
      statusText: response.statusText,
    })
  },
  onError({ error }) {
    console.error('API Error:', error)
    if (!isApiError(error)) {
      handleUnexpectedError(error)
    }
  },
})

export const client = {
  GET: instance.GET,
  POST: instance.POST,
  PUT: instance.PUT,
  DELETE: instance.DELETE,
}
