// @flow
import * as React from 'react'
import {globalStyles} from '../../styles'

import type {Props} from './container'

const Container = ({children, onBack, style, outerStyle}: Props) => {
  return (
    <div style={{...stylesContainer, ...outerStyle}}>
      <div style={{...stylesInnerContainer, ...style}}>{children}</div>
    </div>
  )
}

const stylesContainer = {
  ...globalStyles.flexBoxColumn,
  alignItems: 'flex-start',
  flex: 1,
  justifyContent: 'flex-start',
  padding: 64,
}

const stylesInnerContainer = {
  ...globalStyles.flexBoxColumn,
  alignSelf: 'stretch',
  height: '100%',
  width: '100%',
}

export default Container
