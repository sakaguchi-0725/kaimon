import { Colors } from '@/shared/constants'
import { MemberStatus } from '.'

export const StatusLabels: Record<MemberStatus, string> = {
  active: '参加済',
  pending: '招待済',
}

export const StatusColors: Record<MemberStatus, string> = {
  active: Colors.success,
  pending: Colors.warning,
}
