import { onMounted, ref, watch } from 'vue'
import type { GetGroupResponse } from '../model'
import { client } from '@/shared/api'

export const useGroupInfo = (id: () => string) => {
  const group = ref<GetGroupResponse>()

  const fetchGroupInfo = async () => {
    const { data: groupData } = await client.GET('/groups/{id}', {
      params: {
        path: { id: id() },
      },
    })

    group.value = groupData
  }

  onMounted(fetchGroupInfo)
  watch(id, fetchGroupInfo)

  return { group }
}
