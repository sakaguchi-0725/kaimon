import { ApiSchema } from '@/shared/api'
import { Colors } from '@/shared/constants'

export type JoinedGroup = ApiSchema<'JoinedGroup'>
export type GroupInfo = ApiSchema<'GetGroupResponse'>
export type Member = ApiSchema<'Member'>

// Member関連の型定義（features/memberから移行）
export type MemberRole = 'admin' | 'member'
export type MemberStatus = 'active' | 'pending'

// ステータス表示用の定数
export const StatusLabels: Record<MemberStatus, string> = {
  active: '参加済',
  pending: '招待済',
}

export const StatusColors: Record<MemberStatus, string> = {
  active: Colors.success,
  pending: Colors.warning,
}
