import * as React from "react";

import UpdateUserMutation from "../queries/mUpdateUser";
import RemoveUserMutation from "../queries/mRemoveUser";
import User from "../queries/qUser";
import UserDetailsPage from "../components/UserDetailsPage";
import Navigator from "../../components/Navigator";
import Notificator, { NotificationType } from "../../components/Notificator";
import i18n from "../../i18n";
import { UpdateUser } from "../queries/types/UpdateUser";
import urls from "../../urls";

interface Props {
  id: string;
}

export const UserDetails: React.StatelessComponent<Props> = ({ id }) => (
  <Notificator>
    {notify => (
      <Navigator>
        {navigate => {
          const handleError = () =>
            notify({
              text: i18n.t("Something went wrong", {
                context: "notification",
              }),
              type: NotificationType.ERROR,
            });
          const handleUpdateUser = (data: UpdateUser) => {
            if (
              data &&
              data.updateUser &&
              data.updateUser.errors &&
              data.updateUser.errors.length > 0
            ) {
              handleError();
            } else {
              notify({
                text: i18n.t("Updated user", {
                  context: "notification",
                }),
              });
            }
          };
          const handleRemoveUser = () => {
            notify({
              text: i18n.t("Deleted user", {
                context: "notification",
              }),
            });
            navigate(urls.userList);
          };
          return (
            <User variables={{ id }}>
              {userData => (
                <RemoveUserMutation onCompleted={handleRemoveUser}>
                  {removeUser => (
                    <UpdateUserMutation onCompleted={handleUpdateUser}>
                      {updateUser => (
                        <UserDetailsPage
                          disabled={userData.loading}
                          loading={userData.loading}
                          onBack={() => navigate(urls.userList)}
                          onDelete={() => removeUser({ variables: { id } })}
                          title={
                            userData.data && userData.data.user
                              ? userData.data.user.email
                              : undefined
                          }
                          transaction={userData.loading ? "loading" : "default"}
                          user={userData.data ? userData.data.user : undefined}
                          onSubmit={variables =>
                            updateUser({ variables: { id, input: variables } })
                          }
                        />
                      )}
                    </UpdateUserMutation>
                  )}
                </RemoveUserMutation>
              )}
            </User>
          );
        }}
      </Navigator>
    )}
  </Notificator>
);
export default UserDetails;
