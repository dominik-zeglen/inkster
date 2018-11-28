import * as React from "react";

import CreateUserMutation from "../queries/mCreateUser";
import Users from "../queries/qUsers";
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
          <Users>
            {users => {
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
                      disabled={users.loading || createUserData.loading}
                      loading={users.loading || createUserData.loading}
                      users={users.data ? users.data.users : undefined}
                      onAdd={data => createUser({ variables: { input: data } })}
                      onNextPage={() => undefined}
                      onPreviousPage={() => undefined}
                      onRowClick={id => () => navigate(urls.userDetails(id))}
                    />
                  )}
                </CreateUserMutation>
              );
            }}
          </Users>
        )}
      </Notificator>
    )}
  </Navigator>
);
export default UserList;
