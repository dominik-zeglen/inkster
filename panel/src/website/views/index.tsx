import * as React from "react";
import { Switch, Route } from "react-router-dom";

import Website from "./Website";
import { paths } from "../../urls";

export const PageSection: React.StatelessComponent = () => (
  <Switch>
    <Route path={paths.websiteSettings} component={Website} />
  </Switch>
);
export default PageSection;
