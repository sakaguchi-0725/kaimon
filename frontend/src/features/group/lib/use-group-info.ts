import { onMounted, ref, watch } from 'vue'
import type { GetGroupResponse, Member } from '../model'
import { client } from '@/shared/api'

export const useGroupInfo = (id: () => string) => {
  const group = ref<GetGroupResponse>()
  const members = ref<Member[]>()

  const fetchGroupInfo = async () => {
    const { data: groupData } = await client.GET('/groups/{id}', {
      params: {
        path: { id: id() },
      },
    })

    const { data: membersData } = await client.GET('/groups/{id}/members', {
      params: {
        path: { id: id() },
      },
    })

    group.value = groupData
    members.value = membersData?.members
  }

  onMounted(fetchGroupInfo)
  watch(id, fetchGroupInfo)

  return {
    group,
    members,
  }
}
