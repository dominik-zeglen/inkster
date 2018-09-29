import * as React from "react";
import { Mutation, Query } from "react-apollo";

import mUpdateUser, {
  Result as UpdateUserResult
} from "../queries/mUpdateUser";
import mRemoveUser from "../queries/mRemoveUser";
import qUser from "../queries/qUser";
import UserDetailsPage from "../components/UserDetailsPage";
import Navigator from "../../components/Navigator";
import Notificator, { NotificationType } from "../../components/Notificator";
import i18n from "../../i18n";

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
                context: "notification"
              }),
              type: NotificationType.ERROR
            });
          const handleUpdateUser = (data: UpdateUserResult) => {
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
                  context: "notification"
                })
              });
            }
          };
          const handleRemoveUser = () => {
            notify({
              text: i18n.t("Deleted user", {
                context: "notification"
              })
            });
            navigate("/users/");
          };
          return (
            <Query query={qUser} variables={{ id }}>
              {userData => (
                <Mutation
                  mutation={mRemoveUser}
                  onCompleted={handleRemoveUser}
                  onError={handleError}
                >
                  {removeUser => (
                    <Mutation
                      mutation={mUpdateUser}
                      onCompleted={handleUpdateUser}
                      onError={handleError}
                    >
                      {updateUser => (
                        <UserDetailsPage
                          disabled={userData.loading}
                          loading={userData.loading}
                          onBack={() => navigate("/users/")}
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
                    </Mutation>
                  )}
                </Mutation>
              )}
            </Query>
          );
        }}
      </Navigator>
    )}
  </Notificator>
);
export default UserDetails;
