import {
  createRouter,
  createWebHistory,
  createMemoryHistory,
  type RouteRecordRaw,
} from 'vue-router'
import { h } from 'vue'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'sample',
    component: {
      render() {
        return h('h1', 'Hello World')
      },
    },
  },
]

export const createAppRouter = (type: 'web' | 'memory') => {
  const history = type === 'web' ? createWebHistory() : createMemoryHistory()
  const router = createRouter({ history, routes })

  return router
}

export default routes
