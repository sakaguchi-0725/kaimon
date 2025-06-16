import {
  LoginPage,
  ResetPassword,
  ResetPasswordConfirm,
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
      path: '',
      component: AuthLayout,
      children: [
        {
          path: '/login',
          name: 'login',
          component: LoginPage,
        },
        {
          path: '/signup',
          name: 'signup',
          component: SignupPage,
        },
        {
          path: '/signup-confirm',
          name: 'signup-confirm',
          component: SignupConfirmPage,
        },
        {
          path: '/reset-password',
          name: 'reset-password',
          component: ResetPassword,
        },
        {
          path: '/reset-password-confirm',
          name: 'reset-password-confirm',
          component: ResetPasswordConfirm,
        },
      ],
    },
  ]
}
