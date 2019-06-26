import * as React from "react";
import { BrowserRouter } from "react-router-dom";
import { ThemeProvider as AuroraThemeProvider } from "aurora-ui-kit/dist/utils/jss";

import theme from "../src/theme";
import Baseline from "aurora-ui-kit/dist/components/Baseline";

export const Decorator = (storyFn: any) => (
  <BrowserRouter>
    <AuroraThemeProvider theme={theme}>
      <>
        <Baseline />
        {storyFn()}
      </>
    </AuroraThemeProvider>
  </BrowserRouter>
);
export default Decorator;
