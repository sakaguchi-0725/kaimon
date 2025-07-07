import { View, Text, StyleSheet, TouchableOpacity } from 'react-native'
import { useForm, Controller } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { Colors } from '@/shared/constants'
import { BottomSheetModal } from '@/shared/ui/modal'
import { Button } from '@/shared/ui'
import { TextInput, Textarea } from '@/shared/ui/input'
import { X } from 'react-native-feather'
import { createGroupSchema, type CreateGroupForm } from '../model/schema'
import { useCreateGroup } from '../lib'

type Props = {
  isVisible: boolean
  onClose: () => void
  onSuccess?: () => void
  onRefetch?: () => void
}

export const CreateGroupModal = ({
  isVisible,
  onClose,
  onSuccess,
  onRefetch,
}: Props) => {
  const { createGroup, isLoading, error } = useCreateGroup()

  const {
    control,
    handleSubmit,
    formState: { errors, isValid },
    reset,
  } = useForm<CreateGroupForm>({
    resolver: zodResolver(createGroupSchema),
    defaultValues: {
      name: '',
      description: '',
    },
  })

  const onSubmit = async (data: CreateGroupForm) => {
    try {
      await createGroup(data, () => {
        onRefetch?.()
      })
      reset()
      onSuccess?.()
      onClose()
    } catch (error) {
      console.error('グループ作成エラー:', error)
    }
  }

  const handleClose = () => {
    reset()
    onClose()
  }

  return (
    <BottomSheetModal isVisible={isVisible} onClose={handleClose}>
      <View style={styles.modalHeader}>
        <Text style={styles.title}>グループを作成</Text>
        <TouchableOpacity onPress={handleClose}>
          <X width={24} height={24} stroke={Colors.mainText} />
        </TouchableOpacity>
      </View>

      <View style={styles.formContainer}>
        {/* グループ名入力 */}
        <Controller
          control={control}
          name="name"
          render={({ field: { onChange, value } }) => (
            <TextInput
              label="グループ名"
              required
              value={value}
              onChangeText={onChange}
              placeholder="例: 家族の買い物リスト"
              error={errors.name?.message}
            />
          )}
        />

        {/* 説明入力 */}
        <Controller
          control={control}
          name="description"
          render={({ field: { onChange, value } }) => (
            <Textarea
              label="説明"
              value={value || ''}
              onChangeText={onChange}
              placeholder="例: 家族みんなで使う買い物リスト"
              numberOfLines={3}
              error={errors.description?.message}
            />
          )}
        />

        {/* エラーメッセージ */}
        {error && <Text style={styles.errorText}>{error}</Text>}
      </View>

      {/* 作成ボタン */}
      <Button
        text={isLoading ? '作成中...' : '作成'}
        onPress={handleSubmit(onSubmit)}
        size="full"
        variant="solid"
        color="primary"
        disabled={!isValid || isLoading}
      />
    </BottomSheetModal>
  )
}

const styles = StyleSheet.create({
  modalHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 16,
  },
  title: {
    fontSize: 20,
    fontWeight: 'bold',
    color: Colors.mainText,
  },
  formContainer: {
    gap: 16,
    marginBottom: 18,
  },
  errorText: {
    color: Colors.error,
    fontSize: 14,
    textAlign: 'center',
  },
})
