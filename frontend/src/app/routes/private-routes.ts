import { GroupListPage, GroupDetailPage } from '@/pages/group'
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
        {
          path: ':id',
          name: 'group-detail',
          meta: {
            title: 'グループ詳細',
            isBackButton: true,
            toBack: '/groups',
          },
          component: GroupDetailPage,
        },
      ],
    },
  ]
}
