import * as React from "react";
import { Switch, Route, RouteComponentProps } from "react-router-dom";
import { parse as parseQs } from "qs";

import UserListComponent, {
  QueryParams as UserListQueryParams,
} from "./UserList";
import UserDetailsComponent, {
  QueryParams as UserDetailsQueryParams,
} from "./UserDetails";
import { paths } from "../../urls";

interface UserDetailsRouteParams {
  id: string;
}
const UserDetails: React.StatelessComponent<
  RouteComponentProps<UserDetailsRouteParams>
> = ({ match, location }) => {
  const qs = parseQs(location.search.substr(1));
  const params: UserDetailsQueryParams = {
    modal: qs.modal,
  };
  const decodedId = decodeURIComponent(match.params.id);
  return <UserDetailsComponent id={decodedId} params={params} />;
};

const UserList: React.StatelessComponent<
  RouteComponentProps<UserDetailsRouteParams>
> = ({ location }) => {
  const qs = parseQs(location.search.substr(1));
  const params: UserListQueryParams = {
    modal: qs.modal,
  };
  return <UserListComponent params={params} />;
};

export const UserSection: React.StatelessComponent<
  RouteComponentProps<{}>
> = () => (
  <Switch>
    <Route path={paths.userDetails(":id")} component={UserDetails} />
    <Route path={paths.userList} component={UserList} />
  </Switch>
);
export default UserSection;
