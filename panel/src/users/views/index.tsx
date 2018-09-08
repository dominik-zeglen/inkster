import * as React from "react";
import { Switch, Route, RouteComponentProps } from "react-router-dom";

import UserList from "./UserList";
// import { unurlize } from "../../utils";

export const UserSection: React.StatelessComponent<RouteComponentProps<{}>> = ({
  match
}) => (
  <Switch>
    <Route path={`${match.url}/`} component={UserList} />
  </Switch>
);
export default UserSection;
