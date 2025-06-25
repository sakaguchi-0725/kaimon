import React from 'react'
import { StyleSheet, Text, View } from 'react-native'
import { MemberStatus } from '../model'
import { StatusColors, StatusLabels } from '../model/constants'

type StatusLabelProps = {
  status: MemberStatus
}

export const StatusLabel: React.FC<StatusLabelProps> = ({ status }) => {
  return (
    <View style={styles.statusContainer}>
      <View
        style={[styles.statusDot, { backgroundColor: StatusColors[status] }]}
      />
      <Text style={[styles.statusText, { color: StatusColors[status] }]}>
        {StatusLabels[status]}
      </Text>
    </View>
  )
}

const styles = StyleSheet.create({
  statusContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: 3,
  },
  statusDot: {
    width: 8,
    height: 8,
    borderRadius: 4,
  },
  statusText: {
    fontSize: 14,
  },
})
