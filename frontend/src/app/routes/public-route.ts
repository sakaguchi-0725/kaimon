import {
  LoginPage,
  ResetPasswordPage,
  ResetPasswordConfirmPage,
  SignupConfirmPage,
  SignupPage,
} from '@/pages/auth'
import { WelcomePage } from '@/pages/welcome'
import type { RouteRecordRaw } from 'vue-router'
import { AuthLayout, PublicLayout } from '../layout'

export const getPublicRoutes = (): RouteRecordRaw[] => {
  return [
    {
      path: '/',
      component: PublicLayout,
      children: [
        {
          path: '',
          name: 'welcome',
          component: WelcomePage,
        },
      ],
    },
    {
      path: '/auth',
      component: AuthLayout,
      children: [
        {
          path: 'login',
          name: 'login',
          meta: {
            title: 'ログイン',
          },
          component: LoginPage,
        },
        {
          path: 'signup',
          name: 'signup',
          meta: {
            title: 'アカウント登録',
          },
          component: SignupPage,
        },
        {
          path: 'signup/confirm',
          name: 'signup-confirm',
          meta: {
            title: 'アカウント登録確認',
          },
          component: SignupConfirmPage,
        },
        {
          path: 'reset-password',
          name: 'reset-password',
          meta: {
            title: 'パスワードリセット',
          },
          component: ResetPasswordPage,
        },
        {
          path: 'reset-password/confirm',
          name: 'reset-password-confirm',
          meta: {
            title: 'パスワードリセット確認',
          },
          component: ResetPasswordConfirmPage,
        },
      ],
    },
  ]
}
