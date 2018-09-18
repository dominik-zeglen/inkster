import * as React from "react";
import { Mutation, Query } from "react-apollo";

import mCreateUser, {
  Result as CreateUserResult
} from "../queries/mCreateUser";
import qUsers from "../queries/qUsers";
import UserListPage from "../components/UserListPage";
import Navigator from "../../components/Navigator";
import { urls } from "../../";

export const UserList: React.StatelessComponent<{}> = () => (
  <Navigator>
    {navigate => (
      <Query query={qUsers} fetchPolicy="cache-and-network">
        {({ data, loading, error }) => {
          const handleAddUser = (data: CreateUserResult) => {
            if (data.createUser.errors.length === 0) {
              navigate(urls.userDetails(data.createUser.user.id));
            }
          };
          return (
            <Mutation mutation={mCreateUser} onCompleted={handleAddUser}>
              {(createUser, createUserData) => (
                <UserListPage
                  disabled={loading || createUserData.loading}
                  loading={loading || createUserData.loading}
                  users={data ? data.users : undefined}
                  onAdd={data => createUser({ variables: { input: data } })}
                  onNextPage={() => undefined}
                  onPreviousPage={() => undefined}
                  onRowClick={id => () => navigate(urls.userDetails(id))}
                />
              )}
            </Mutation>
          );
        }}
      </Query>
    )}
  </Navigator>
);
export default UserList;
