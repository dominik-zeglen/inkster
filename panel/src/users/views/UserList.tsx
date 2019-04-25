import * as React from "react";
import Input from "aurora-ui-kit/dist/components/TextInput";

import CreateUserMutation from "../queries/mCreateUser";
import Users from "../queries/qUsers";
import UserListPage from "../components/UserListPage";
import Navigator from "../../components/Navigator";
import Notificator, { NotificationType } from "../../components/Notificator";
import urls from "../../urls";
import i18n from "../../i18n";
import { CreateUser } from "../queries/types/CreateUser";
import { Modal, Pagination } from "../../types";
import FormDialog from "../../components/FormDialog";
import { mergeQs, maybe } from "../../utils";
import Paginator, { createPaginationState } from "../../components/Paginator";
import { PAGINATE_BY } from "../..";

export type QueryParams = Partial<Modal<"create-user"> & Pagination>;
export interface Props {
  params: QueryParams;
}

export const UserList: React.StatelessComponent<Props> = ({ params }) => (
  <Navigator>
    {navigate => (
      <Notificator>
        {notify => {
          const paginationState = createPaginationState(PAGINATE_BY, params);
          return (
            <Users variables={{ paginate: paginationState }}>
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
                        <Paginator
                          pageInfo={maybe(() => users.data.users.pageInfo)}
                          paginationState={paginationState}
                          queryString={params}
                        >
                          {({ loadNextPage, loadPreviousPage, pageInfo }) => (
                            <UserListPage
                              disabled={users.loading || createUserData.loading}
                              loading={users.loading || createUserData.loading}
                              users={maybe(() =>
                                users.data.users.edges.map(edge => edge.node),
                              )}
                              pageInfo={pageInfo}
                              onAdd={() =>
                                navigate(
                                  mergeQs(params, {
                                    modal: "create-user",
                                  }),
                                )
                              }
                              onNextPage={loadNextPage}
                              onPreviousPage={loadPreviousPage}
                              onRowClick={id => () =>
                                navigate(urls.userDetails(id))}
                            />
                          )}
                        </Paginator>
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
                              onChange={email =>
                                change({
                                  target: {
                                    name: "email",
                                    value: email,
                                  },
                                } as any)
                              }
                              InputProps={{
                                componentProps: {
                                  type: "email",
                                },
                              }}
                              value={data.email}
                              label={i18n.t("User email")}
                            />
                          )}
                        </FormDialog>
                      </>
                    )}
                  </CreateUserMutation>
                );
              }}
            </Users>
          );
        }}
      </Notificator>
    )}
  </Navigator>
);
export default UserList;
