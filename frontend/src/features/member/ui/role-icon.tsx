import { User, Users } from 'react-native-feather'
import { MemberRole } from '../model'
import { Colors } from '@/shared/constants'

type Props = {
  role: MemberRole
}

export const MemberRoleIcon = ({ role }: Props) => {
  if (role === 'admin') return <Users stroke={Colors.primary} />
  if (role === 'member') return <User stroke={Colors.primary} />
}
