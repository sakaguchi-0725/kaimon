import { useForm } from 'vee-validate'
import { createGroupSchema, type CreateGroupForm } from '../model'
import { toTypedSchema } from '@vee-validate/zod'
import { client } from '@/shared/api'
import { showSuccess } from '@/shared/ui/toast'

export const useCreateGroup = () => {
  const { defineField, errors, handleSubmit } = useForm<CreateGroupForm>({
    validationSchema: toTypedSchema(createGroupSchema),
  })

  const onSubmit = (onSuccess: () => void) => {
    return handleSubmit(async values => {
      const { error } = await client.POST('/groups', {
        body: {
          name: values.name,
          description: values.description,
        },
      })

      if (error) return

      showSuccess('作成完了', 'グループを作成しました')
      onSuccess()
    })
  }

  return {
    defineField,
    errors,
    onSubmit,
  }
}
