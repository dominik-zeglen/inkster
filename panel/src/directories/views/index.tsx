import * as React from "react";
import { Switch, Route } from "react-router-dom";

import DirectoryRoot from "./DirectoryRoot";

export const DirectorySection = () => (
  <Switch>
    <Route exact path="/" component={DirectoryRoot} />
  </Switch>
);
export default DirectorySection;
