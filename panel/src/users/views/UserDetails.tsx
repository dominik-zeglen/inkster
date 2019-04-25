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
import { Modal, Pagination } from "../../types";
import ActionDialog from "../../components/ActionDialog";
import { mergeQs, maybe } from "../../utils";
import Paginator, { createPaginationState } from "../../components/Paginator";

export type QueryParams = Partial<Modal<"remove">> & Pagination;
export interface Props {
  id: string;
  params: QueryParams;
}

const PAGINATE_BY = 5;

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

          const paginationState = createPaginationState(PAGINATE_BY, params);

          return (
            <User variables={{ id, paginate: paginationState }}>
              {userData => (
                <RemoveUserMutation onCompleted={handleRemoveUser}>
                  {removeUser => (
                    <UpdateUserMutation onCompleted={handleUpdateUser}>
                      {updateUser => (
                        <Paginator
                          pageInfo={maybe(
                            () => userData.data.user.pages.pageInfo,
                          )}
                          paginationState={paginationState}
                          queryString={params}
                        >
                          {({ loadNextPage, loadPreviousPage, pageInfo }) => (
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
                                title={maybe(() => userData.data.user.email)}
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
                                onNextPage={loadNextPage}
                                onPreviousPage={loadPreviousPage}
                                pageInfo={pageInfo}
                                onRowClick={pageId => () =>
                                  navigate(urls.pageDetails(pageId))}
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
                                onConfirm={() =>
                                  removeUser({ variables: { id } })
                                }
                              >
                                {i18n.t(
                                  "Are you sure you want to remove {{ email }}?",
                                  {
                                    email: maybe(
                                      () => userData.data.user.email,
                                    ),
                                  },
                                )}
                              </ActionDialog>
                            </>
                          )}
                        </Paginator>
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
