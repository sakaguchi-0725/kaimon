import {
  StyleSheet,
  Text,
  View,
  Dimensions,
  Image,
  FlatList,
} from 'react-native'
import { useNavigation } from '@react-navigation/native'
import { AuthNavigationProp } from '@/screens/auth'
import { useRef, useState, useEffect } from 'react'
import { Colors } from '@/shared/constants'
import { Button } from '@/shared/ui'

const { width } = Dimensions.get('window')

// カルーセルに表示する画像の配列
const carouselImages = [
  'https://placehold.jp/fb5300/ffff/400x600.png?text=%E3%83%80%E3%83%9F%E3%83%BC1',
  'https://placehold.jp/fb5300/ffff/400x600.png?text=%E3%83%80%E3%83%9F%E3%83%BC2',
  'https://placehold.jp/fb5300/ffff/400x600.png?text=%E3%83%80%E3%83%9F%E3%83%BC3',
]

export const WelcomeScreen = () => {
  const navigation = useNavigation<AuthNavigationProp>()
  const flatListRef = useRef<FlatList>(null)
  const [currentIndex, setCurrentIndex] = useState(0)

  // 自動スクロール
  useEffect(() => {
    const interval = setInterval(() => {
      const nextIndex = (currentIndex + 1) % carouselImages.length
      flatListRef.current?.scrollToIndex({
        index: nextIndex,
        animated: true,
      })
      setCurrentIndex(nextIndex)
    }, 3000)

    return () => clearInterval(interval)
  }, [currentIndex])

  return (
    <View style={styles.container}>
      {/* 画面上部の画像スライダー (70%) */}
      <View style={styles.carouselContainer}>
        <FlatList
          ref={flatListRef}
          data={carouselImages}
          horizontal
          pagingEnabled
          showsHorizontalScrollIndicator={false}
          onMomentumScrollEnd={(event) => {
            const slideIndex = Math.floor(
              event.nativeEvent.contentOffset.x / width,
            )
            setCurrentIndex(slideIndex)
          }}
          renderItem={({ item }) => (
            <View style={[styles.carouselItem, { width }]}>
              <Image
                source={{ uri: item }}
                style={styles.carouselImage}
                resizeMode="cover"
              />
            </View>
          )}
          keyExtractor={(_, index) => index.toString()}
        />
      </View>

      {/* 画面下部のコンテンツ (30%) */}
      <View style={styles.contentContainer}>
        <Text style={styles.subtitle}>
          家族や友人と買い物リストを共有しましょう
        </Text>

        {/* ページインジケーター */}
        <View style={styles.pagination}>
          {carouselImages.map((_, index) => (
            <View
              key={index}
              style={[
                styles.paginationDot,
                index === currentIndex && styles.paginationDotActive,
              ]}
            />
          ))}
        </View>

        <View style={styles.buttonContainer}>
          <Button
            text="新規登録"
            onPress={() => navigation.navigate('SignUp')}
            size="full"
            variant="solid"
            color="primary"
          />

          <Button
            text="ログイン"
            onPress={() => navigation.navigate('Login')}
            size="full"
            variant="text"
            color="primary"
          />
        </View>
      </View>
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
  },
  carouselContainer: {
    height: '60%',
  },
  carouselItem: {
    justifyContent: 'center',
    alignItems: 'center',
  },
  carouselImage: {
    width: '100%',
    height: '100%',
  },
  pagination: {
    width: '100%',
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    marginBottom: 30,
  },
  paginationDot: {
    width: 8,
    height: 8,
    borderRadius: 4,
    backgroundColor: 'rgba(0, 0, 0, 0.3)',
    margin: 5,
  },
  paginationDotActive: {
    backgroundColor: Colors.primary,
  },
  contentContainer: {
    flex: 1,
    paddingHorizontal: 20,
    justifyContent: 'center',
    alignItems: 'center',
  },
  subtitle: {
    fontSize: 16,
    textAlign: 'center',
    marginBottom: 20,
    color: Colors.mainText,
  },
  buttonContainer: {
    width: '100%',
    gap: 16,
  },
})
