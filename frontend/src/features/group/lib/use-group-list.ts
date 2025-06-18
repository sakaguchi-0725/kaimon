import { computed, onMounted, ref } from 'vue'
import type { JoinedGroupsResponse } from '../model'
import { client } from '@/shared/api'
import { useRouter } from 'vue-router'

export const useGroupList = () => {
  const router = useRouter()
  const groups = ref<JoinedGroupsResponse['groups']>([])

  onMounted(async () => {
    const { data } = await client.GET('/groups')
    groups.value = data?.groups ?? []
  })

  const hasGroups = computed(() => groups.value.length > 0)

  const onDetail = (id: string) => {
    router.push({ name: 'group-detail', params: { id } })
  }

  return {
    groups,
    hasGroups,
    onDetail,
  }
}
