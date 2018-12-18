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
import { Modal } from "../../types";
import ActionDialog from "../../components/ActionDialog";
import { mergeQs, maybe } from "../../utils";

export type QueryParams = Partial<Modal<"remove">>;
export interface Props {
  id: string;
  params: QueryParams;
}

export const UserDetails: React.StatelessComponent<Props> = ({
  id,
  params,
}) => (
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
                        <>
                          <UserDetailsPage
                            disabled={userData.loading}
                            loading={userData.loading}
                            onBack={() => navigate(urls.userList)}
                            onDelete={() =>
                              navigate(
                                mergeQs(params, {
                                  modal: "remove",
                                }),
                              )
                            }
                            title={
                              userData.data && userData.data.user
                                ? userData.data.user.email
                                : undefined
                            }
                            transaction={
                              userData.loading ? "loading" : "default"
                            }
                            user={
                              userData.data ? userData.data.user : undefined
                            }
                            onSubmit={variables =>
                              updateUser({
                                variables: { id, input: variables },
                              })
                            }
                          />
                          <ActionDialog
                            show={params.modal === "remove"}
                            size="xs"
                            title={i18n.t("Remove user")}
                            onClose={() =>
                              navigate(
                                mergeQs(params, {
                                  modal: undefined,
                                }),
                              )
                            }
                            onConfirm={() => removeUser({ variables: { id } })}
                          >
                            {i18n.t(
                              "Are you sure you want to remove {{ email }}?",
                              {
                                email: maybe(() => userData.data.user.email),
                              },
                            )}
                          </ActionDialog>
                        </>
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
