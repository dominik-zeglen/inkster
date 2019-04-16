import * as React from "react";
import { ThemeProvider } from "react-jss";
import { BrowserRouter } from "react-router-dom";
import auroraTheme from "aurora-ui-kit/dist/theme";
import { ThemeProvider as AuroraThemeProvider } from "aurora-ui-kit/dist/utils/styled-components";

import GlobalStylesheet from "../src/Stylesheet";
import theme from "../src/theme";

export const Decorator = (storyFn: any) => (
  <BrowserRouter>
    <ThemeProvider theme={theme}>
      <AuroraThemeProvider theme={auroraTheme}>
        <>
          <GlobalStylesheet />
          {storyFn()}
        </>
      </AuroraThemeProvider>
    </ThemeProvider>
  </BrowserRouter>
);
export default Decorator;
