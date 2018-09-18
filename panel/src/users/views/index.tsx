import * as React from "react";
import { Switch, Route, RouteComponentProps } from "react-router-dom";

import UserList from "./UserList";
import UserDetailsComponent from './UserDetails'
import { unurlize } from "../../utils";

interface UserDetailsRouteParams {
  id: string
}
const UserDetails: React.StatelessComponent<RouteComponentProps<UserDetailsRouteParams>> = ({ match }) => {
  const decodedId = unurlize(match.params.id)
  return <UserDetailsComponent id={decodedId} />
}

export const UserSection: React.StatelessComponent<RouteComponentProps<{}>> = ({
  match
}) => (
  <Switch>
    <Route path={`${match.url}/:id/`} component={UserDetails} />
    <Route path={`${match.url}/`} component={UserList} />
  </Switch>
);
export default UserSection;
