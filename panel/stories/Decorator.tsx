import * as React from "react";
import { ThemeProvider } from "react-jss";
import { BrowserRouter } from "react-router-dom";

import GlobalStylesheet from "../src/Stylesheet";
import theme from "../src/theme";

export const Decorator = (storyFn: any) => (
  <BrowserRouter>
    <ThemeProvider theme={theme}>
      <>
        <GlobalStylesheet />
        {storyFn()}
      </>
    </ThemeProvider>
  </BrowserRouter>
);
export default Decorator;
