import {
  createRouter,
  createWebHistory,
  createMemoryHistory,
  type RouteRecordRaw,
} from 'vue-router'
import { getPublicRoutes } from './public-route'
import { getPrivateRoutes } from './private-routes'

const routes: RouteRecordRaw[] = [...getPublicRoutes(), ...getPrivateRoutes()]

export const createAppRouter = (type: 'web' | 'memory') => {
  const history = type === 'web' ? createWebHistory() : createMemoryHistory()
  const router = createRouter({ history, routes })

  return router
}

export default routes
