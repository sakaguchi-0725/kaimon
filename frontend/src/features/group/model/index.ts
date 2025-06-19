import { type ApiSchema } from '@/shared/api'
import { createGroupSchema } from './schema'
import type { infer as Infer } from 'zod'

export * from './schema'

export type CreateGroupForm = Infer<typeof createGroupSchema>

export type JoinedGroupsResponse = ApiSchema<'GetJoinedGroupsResponse'>
export type GetGroupResponse = ApiSchema<'GetGroupResponse'>
