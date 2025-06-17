import { GroupListPage } from '@/pages/group-list'
import type { RouteRecordRaw } from 'vue-router'
import { PrivateLayout } from '../layout'

export const getPrivateRoutes = (): RouteRecordRaw[] => {
  return [
    {
      path: '/groups',
      component: PrivateLayout,
      children: [
        {
          path: '',
          name: 'group-list',
          meta: {
            title: 'グループ一覧',
            isBackButton: false,
          },
          component: GroupListPage,
        },
      ],
    },
  ]
}
