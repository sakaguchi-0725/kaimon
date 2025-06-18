import { onMounted, ref, watch } from 'vue'
import type { ShoppingItem } from '../model'
import { client } from '@/shared/api'

export const useGetShoppingItem = (groupId: () => string) => {
  const items = ref<ShoppingItem[]>()

  const fetchItems = async () => {
    const { data } = await client.GET('/groups/{id}/items', {
      params: {
        path: { id: groupId() },
      },
    })

    items.value = data?.items
  }

  onMounted(fetchItems)
  watch(groupId, fetchItems)

  return { items }
}
