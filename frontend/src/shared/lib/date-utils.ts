export const formatDate = (dateString?: string): string => {
  if (!dateString) return ''

  try {
    const date = new Date(dateString)

    // 無効な日付の場合
    if (isNaN(date.getTime())) {
      return dateString
    }

    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, '0')
    const day = String(date.getDate()).padStart(2, '0')

    return `${year}/${month}/${day}`
  } catch (error) {
    // パースエラーの場合は元の文字列を返す
    return dateString
  }
}
