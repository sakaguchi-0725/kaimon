export type MemberRole = 'admin' | 'member'

export type MemberStatus = 'active' | 'pending'

export type Member = {
  id: string
  name: string
  role: MemberRole
  status: MemberStatus
  joinedAt?: string
}
