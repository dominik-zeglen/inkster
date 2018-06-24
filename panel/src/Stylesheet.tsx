import * as React from 'react'
import withStyles from 'react-jss'

import stylesheet from './stylesheet'

export const GlobalStylesheet = withStyles(stylesheet)(() => <div />);
export default GlobalStylesheet;
