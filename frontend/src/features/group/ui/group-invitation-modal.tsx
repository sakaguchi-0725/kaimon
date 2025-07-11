import React from 'react'
import {
  View,
  Text,
  StyleSheet,
  TouchableOpacity,
  TextInput,
} from 'react-native'
import { Colors } from '@/shared/constants'
import { BottomSheetModal } from '@/shared/ui/modal'
import { X, Copy } from 'react-native-feather'

type Props = {
  isVisible: boolean
  onClose: () => void
  invitationCode?: string
}

export const GroupInvitationModal = ({
  isVisible,
  onClose,
  invitationCode = 'ABC123XY', // 仮の招待コード
}: Props) => {
  const handleCopyCode = () => {
    console.log('招待コードをコピー:', invitationCode)
    // TODO: クリップボードにコピーする処理を実装
  }

  return (
    <BottomSheetModal isVisible={isVisible} onClose={onClose}>
      <View style={styles.modalHeader}>
        <Text style={styles.title}>メンバーを招待</Text>
        <TouchableOpacity
          onPress={onClose}
          accessible={true}
          accessibilityLabel="モーダルを閉じる"
          accessibilityRole="button"
        >
          <X width={24} height={24} stroke={Colors.mainText} />
        </TouchableOpacity>
      </View>

      <View style={styles.content}>
        <Text style={styles.description}>
          招待コードを共有して、メンバーを招待しましょう。
        </Text>

        <View style={styles.codeContainer}>
          <Text style={styles.codeLabel}>招待コード</Text>
          <View style={styles.codeInputContainer}>
            <TextInput
              style={styles.codeInput}
              value={invitationCode}
              editable={false}
              selectTextOnFocus={true}
            />
            <TouchableOpacity
              style={styles.copyButton}
              onPress={handleCopyCode}
              accessible={true}
              accessibilityLabel="招待コードをコピー"
              accessibilityRole="button"
            >
              <Copy width={20} height={20} stroke={Colors.primary} />
            </TouchableOpacity>
          </View>
        </View>

        <Text style={styles.note}>
          招待コードの有効期限は7日間です。{'\n'}
          期限が切れた場合は新しいコードを生成してください。
        </Text>
      </View>
    </BottomSheetModal>
  )
}

const styles = StyleSheet.create({
  modalHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 20,
  },
  title: {
    fontSize: 20,
    fontWeight: 'bold',
    color: Colors.mainText,
  },
  content: {
    gap: 16,
  },
  description: {
    fontSize: 16,
    color: Colors.mainText,
    lineHeight: 22,
  },
  codeContainer: {
    gap: 8,
  },
  codeLabel: {
    fontSize: 14,
    fontWeight: '600',
    color: Colors.mainText,
  },
  codeInputContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    backgroundColor: Colors.white,
    borderWidth: 1,
    borderColor: Colors.border,
    borderRadius: 8,
    paddingHorizontal: 12,
    paddingVertical: 2,
  },
  codeInput: {
    flex: 1,
    fontSize: 18,
    fontWeight: '600',
    color: Colors.mainText,
    paddingVertical: 12,
    textAlign: 'center',
    letterSpacing: 2,
  },
  copyButton: {
    padding: 8,
    marginLeft: 8,
  },
  note: {
    fontSize: 14,
    color: Colors.subText,
    lineHeight: 20,
    textAlign: 'left',
    marginTop: 8,
  },
})
