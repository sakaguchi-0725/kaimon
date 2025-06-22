import { ReactNode, useEffect, useRef } from 'react'
import { Animated, Modal, StyleSheet, TouchableOpacity, View, ViewStyle } from 'react-native'
import { Colors } from '@/shared/constants'

export interface BottomSheetModalProps {
  isVisible: boolean
  onClose: () => void
  children: ReactNode
  containerStyle?: ViewStyle
}

export const BottomSheetModal = ({ 
  isVisible, 
  onClose, 
  children,
  containerStyle
}: BottomSheetModalProps) => {
  const overlayAnimation = useRef(new Animated.Value(0)).current
  const modalAnimation = useRef(new Animated.Value(0)).current

  // モーダルの表示状態が変わったときのアニメーション処理
  useEffect(() => {
    if (isVisible) {
      // モーダルを開く
      // まずオーバーレイをフェードイン
      Animated.timing(overlayAnimation, {
        toValue: 1,
        duration: 80,
        useNativeDriver: true,
      }).start(() => {
        // オーバーレイが表示された後にモーダルをスライドイン
        Animated.timing(modalAnimation, {
          toValue: 1,
          duration: 200,
          useNativeDriver: true,
        }).start()
      })
    } else {
      // モーダルを閉じる
      // まずモーダルをスライドアウト
      Animated.timing(modalAnimation, {
        toValue: 0,
        duration: 200,
        useNativeDriver: true,
      }).start(() => {
        // モーダルが隠れた後にオーバーレイをフェードアウト
        Animated.timing(overlayAnimation, {
          toValue: 0,
          duration: 200,
          useNativeDriver: true,
        }).start()
      })
    }
  }, [isVisible, modalAnimation, overlayAnimation])

  // モーダルのアニメーション
  useEffect(() => {
    if (!isVisible) {
      // モーダルが非表示になったらアニメーション値をリセット
      overlayAnimation.setValue(0)
      modalAnimation.setValue(0)
    }
  }, [isVisible, overlayAnimation, modalAnimation])

  return (
    <Modal
      animationType="none"
      transparent={true}
      visible={isVisible}
      onRequestClose={onClose}
    >
      <View style={styles.modalContainer}>
        {/* 背景オーバーレイ */}
        <Animated.View 
          style={[
            styles.modalOverlay,
            {
              opacity: overlayAnimation
            }
          ]}
          onTouchEnd={onClose}
        />
        
        {/* モーダルコンテンツ */}
        <Animated.View 
          style={[
            styles.modalContent,
            {
              transform: [
                {
                  translateY: modalAnimation.interpolate({
                    inputRange: [0, 1],
                    outputRange: [600, 0]  // スライド距離を大きくして視覚効果を強化
                  })
                }
              ]
            },
            containerStyle
          ]}
        >
          {children}
        </Animated.View>
      </View>
    </Modal>
  )
}

const styles = StyleSheet.create({
  modalContainer: {
    flex: 1,
    justifyContent: 'flex-end',
  },
  modalOverlay: {
    position: 'absolute',
    top: 0,
    left: 0,
    right: 0,
    bottom: 0,
    backgroundColor: 'rgba(0, 0, 0, 0.5)',
  },
  modalContent: {
    backgroundColor: Colors.white,
    borderTopLeftRadius: 16,
    borderTopRightRadius: 16,
    padding: 20,
    maxHeight: '80%',
  },
}) 