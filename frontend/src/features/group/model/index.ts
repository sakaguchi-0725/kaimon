import { type ApiSchema } from '@/shared/api'

export type JoinedGroupsResponse = ApiSchema<'GetJoinedGroupsResponse'>
export type GetGroupResponse = ApiSchema<'GetGroupResponse'>
export type GetGroupMembersResponse = ApiSchema<'GetGroupMembersResponse'>

export type Member = ApiSchema<'Member'>
export type MemberStatus = ApiSchema<'MemberStatus'>
