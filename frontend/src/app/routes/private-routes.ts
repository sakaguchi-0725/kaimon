import { GroupListPage } from '@/pages/group-list'
import type { RouteRecordRaw } from 'vue-router'

export const getPrivateRoutes = (): RouteRecordRaw[] => {
  return [
    {
      path: '/group-list',
      name: 'group-list',
      component: GroupListPage,
    },
  ]
}
