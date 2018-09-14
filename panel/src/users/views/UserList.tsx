import * as React from "react";
import { Mutation, Query } from "react-apollo";

import mCreateUser from "../queries/mCreateUser";
import qUsers from "../queries/qUsers";
import UserListPage from "../components/UserListPage";
import Navigator from "../../components/Navigator";

export const UserList: React.StatelessComponent<{}> = () => (
  <Navigator>
    {navigate => (
      <Query query={qUsers}>
        {({ data, loading, error }) => (
          <Mutation mutation={mCreateUser}>
            {(createUser, createUserData) => (
              <UserListPage
                disabled={loading || createUserData.loading}
                loading={loading || createUserData.loading}
                users={data ? data.users : undefined}
                onAdd={data =>
                  createUser({ variables: { input: data } })
                }
                onNextPage={() => undefined}
                onPreviousPage={() => undefined}
                onRowClick={() => () => undefined}
              />
            )}
          </Mutation>
        )}
      </Query>
    )}
  </Navigator>
);
export default UserList;
