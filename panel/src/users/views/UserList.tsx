import * as React from "react";
import { Query } from "react-apollo";

import CreateUserMutation from "../queries/mCreateUser";
import qUsers from "../queries/qUsers";
import UserListPage from "../components/UserListPage";
import Navigator from "../../components/Navigator";
import Notificator, { NotificationType } from "../../components/Notificator";
import urls from "../../urls";
import i18n from "../../i18n";
import { CreateUser } from "../queries/types/CreateUser";

export const UserList: React.StatelessComponent<{}> = () => (
  <Navigator>
    {navigate => (
      <Notificator>
        {notify => (
          <Query query={qUsers} fetchPolicy="cache-and-network">
            {query => {
              const handleAddUser = (data: CreateUser) => {
                if (data.createUser.errors.length === 0) {
                  notify({
                    text: i18n.t("Sent invitation e-mail", {
                      context: "notification",
                    }),
                  });
                  navigate(urls.userDetails(data.createUser.user.id));
                } else {
                  notify({
                    text: i18n.t("Something went wrong", {
                      context: "notification",
                    }),
                    type: NotificationType.ERROR,
                  });
                }
              };
              return (
                <CreateUserMutation onCompleted={handleAddUser}>
                  {(createUser, createUserData) => (
                    <UserListPage
                      disabled={query.loading || createUserData.loading}
                      loading={query.loading || createUserData.loading}
                      users={query.data ? query.data.users : undefined}
                      onAdd={data => createUser({ variables: { input: data } })}
                      onNextPage={() => undefined}
                      onPreviousPage={() => undefined}
                      onRowClick={id => () => navigate(urls.userDetails(id))}
                    />
                  )}
                </CreateUserMutation>
              );
            }}
          </Query>
        )}
      </Notificator>
    )}
  </Navigator>
);
export default UserList;
