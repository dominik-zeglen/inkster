import * as React from "react";
import { ThemeProvider } from "react-jss";

import GlobalStylesheet from "../src/Stylesheet";
import theme from "../src/theme";

export const Decorator = (storyFn: any) => (
  <ThemeProvider theme={theme}>
    <>
      <GlobalStylesheet />
      {storyFn()}
    </>
  </ThemeProvider>
);
export default Decorator;
