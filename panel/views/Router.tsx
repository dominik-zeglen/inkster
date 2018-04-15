import * as React from "react";
import { Route, Switch } from "react-router";

import ContainerSection from "./containers";

export const Router: React.StatelessComponent = () => (
  <Switch>
    <Route path={"/containers/"} component={ContainerSection} />
  </Switch>
);
export default Router;
