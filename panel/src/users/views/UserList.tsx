import * as React from "react";

import CreateUserMutation from "../queries/mCreateUser";
import Users from "../queries/qUsers";
import UserListPage from "../components/UserListPage";
import Navigator from "../../components/Navigator";
import Notificator, { NotificationType } from "../../components/Notificator";
import urls from "../../urls";
import i18n from "../../i18n";
import { CreateUser } from "../queries/types/CreateUser";
import { Modal } from "../../types";
import FormDialog from "../../components/FormDialog";
import { mergeQs } from "../../utils";
import Input from "../../components/Input";

export type QueryParams = Partial<Modal<"create-user">>;
export interface Props {
  params: QueryParams;
}

export const UserList: React.StatelessComponent<Props> = ({ params }) => (
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
                    <>
                      <UserListPage
                        disabled={users.loading || createUserData.loading}
                        loading={users.loading || createUserData.loading}
                        users={users.data ? users.data.users : undefined}
                        onAdd={() =>
                          navigate(
                            mergeQs(params, {
                              modal: "create-user",
                            }),
                          )
                        }
                        onNextPage={() => undefined}
                        onPreviousPage={() => undefined}
                        onRowClick={id => () => navigate(urls.userDetails(id))}
                      />
                      <FormDialog
                        show={params.modal === "create-user"}
                        width="xs"
                        onClose={() =>
                          navigate(
                            mergeQs(params, {
                              modal: undefined,
                            }),
                          )
                        }
                        onConfirm={data =>
                          createUser({ variables: { input: data } })
                        }
                        title={i18n.t("Add new user")}
                        initial={{ email: "" }}
                      >
                        {({ change, data }) => (
                          <Input
                            name="email"
                            onChange={change}
                            value={data.email}
                            label={i18n.t("User email")}
                            type="email"
                          />
                        )}
                      </FormDialog>
                    </>
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
