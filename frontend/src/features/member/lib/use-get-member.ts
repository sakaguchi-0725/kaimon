import { onMounted, ref, watch } from 'vue'
import type { Member } from '../model'
import { client } from '@/shared/api'

export const useGetMember = (groupId: () => string) => {
  const members = ref<Member[]>()

  const fetchMembers = async () => {
    const { data } = await client.GET('/groups/{id}/members', {
      params: {
        path: { id: groupId() },
      },
    })

    members.value = data?.members
  }

  onMounted(fetchMembers)
  watch(groupId, fetchMembers)

  return { members }
}
