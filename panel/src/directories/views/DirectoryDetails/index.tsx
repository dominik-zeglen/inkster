import * as React from "react";

import MutationProvider from "./MutationProvider";
import Directory from "../../queries/qDirectory";
import DirectoryDetailsPage from "../../components/DirectoryDetailsPage";
import Navigator from "../../../components/Navigator";
import Notificator from "../../../components/Notificator";
import urls from "../../../urls";
import i18n from "../../../i18n";
import { TransactionState, PAGINATE_BY } from "../../../";
import { maybe, mergeQs } from "../../../utils";
import { Modal, Pagination } from "../../../types";
import ActionDialog from "../../../components/ActionDialog";
import FormDialog from "../../../components/FormDialog";
import Input from "../../../components/Input";
import { PageCreate } from "../../queries/types/PageCreate";
import Paginator, {
  createPaginationState,
} from "../../../components/Paginator";

export type QueryParams = Partial<Modal<"remove" | "create-page"> & Pagination>;
interface Props {
  id: string;
  params: QueryParams;
}
interface State {
  transaction: TransactionState;
}
export class DirectoryDetails extends React.Component<Props, State> {
  state = {
    transaction: "default" as "default",
  };

  handleUpdate = () => {
    this.setState({ transaction: "success" });
    setTimeout(() => this.setState({ transaction: "default" }), 3000);
  };

  render() {
    const { id, params } = this.props;
    return (
      <Notificator>
        {notify => (
          <Navigator>
            {navigate => {
              const handleAddPageSuccess = (data: PageCreate) => {
                if (data.createPage.errors.length === 0) {
                  notify({
                    text: i18n.t("Created page", {
                      context: "notification",
                    }),
                  });
                  navigate(urls.pageDetails(data.createPage.page.id));
                }
              };
              const handleRowClick = (pageId: string) => () =>
                navigate(urls.pageDetails(pageId));
              const handleDelete = () => {
                notify({
                  text: i18n.t("Deleted directory", {
                    context: "notification",
                  }),
                });
                navigate(urls.directoryList, true);
              };
              const paginationState = createPaginationState(
                PAGINATE_BY,
                params,
              );
              return (
                <Directory variables={{ id, paginate: paginationState }}>
                  {directory => (
                    <MutationProvider
                      id={id}
                      onDirectoryUpdate={this.handleUpdate}
                      onDirectoryDelete={handleDelete}
                      onPageCreate={handleAddPageSuccess}
                    >
                      {({ createPage, deleteDirectory, updateDirectory }) => (
                        <>
                          <Paginator
                            pageInfo={maybe(
                              () => directory.data.getDirectory.pages.pageInfo,
                            )}
                            paginationState={paginationState}
                            queryString={params}
                          >
                            {({ loadNextPage, loadPreviousPage, pageInfo }) => (
                              <DirectoryDetailsPage
                                directory={maybe(
                                  () => directory.data.getDirectory,
                                )}
                                disabled={directory.loading}
                                loading={directory.loading}
                                transaction={
                                  updateDirectory.opts.loading
                                    ? "loading"
                                    : this.state.transaction
                                }
                                pageInfo={pageInfo}
                                onAdd={() =>
                                  navigate(
                                    mergeQs(params, {
                                      modal: "create-page",
                                    }),
                                  )
                                }
                                onBack={
                                  directory.data && directory.data.getDirectory
                                    ? directory.data.getDirectory.parent &&
                                      directory.data.getDirectory.parent.id
                                      ? () =>
                                          navigate(
                                            urls.directoryDetails(
                                              directory.data.getDirectory.parent
                                                .id,
                                            ),
                                          )
                                      : () => navigate(urls.directoryList)
                                    : () => window.history.back()
                                }
                                onDelete={() =>
                                  navigate(
                                    mergeQs(params, {
                                      modal: "remove",
                                    }),
                                  )
                                }
                                onNextPage={loadNextPage}
                                onPreviousPage={loadPreviousPage}
                                onRowClick={handleRowClick}
                                onSubmit={formData =>
                                  updateDirectory.mutate({
                                    ...formData,
                                    id,
                                  })
                                }
                              />
                            )}
                          </Paginator>
                          <ActionDialog
                            show={params.modal === "remove"}
                            size="xs"
                            title={i18n.t("Remove directory")}
                            onClose={() =>
                              navigate(
                                mergeQs(params, {
                                  modal: undefined,
                                }),
                              )
                            }
                            onConfirm={deleteDirectory.mutate}
                          >
                            {i18n.t(
                              "Are you sure you want to remove {{ name }}?",
                              {
                                name: maybe(
                                  () => directory.data.getDirectory.name,
                                ),
                              },
                            )}
                          </ActionDialog>
                          <FormDialog
                            onClose={() =>
                              navigate(
                                mergeQs(params, {
                                  modal: undefined,
                                }),
                              )
                            }
                            onConfirm={variables =>
                              createPage.mutate({
                                ...variables,
                                parentId: id,
                              })
                            }
                            show={params.modal === "create-page"}
                            title={i18n.t("Create new page")}
                            width="xs"
                            initial={{ name: "" }}
                          >
                            {({ change, data: formData }) => (
                              <Input
                                name="name"
                                onChange={change}
                                value={formData.name}
                                label={i18n.t("Page name")}
                              />
                            )}
                          </FormDialog>
                        </>
                      )}
                    </MutationProvider>
                  )}
                </Directory>
              );
            }}
          </Navigator>
        )}
      </Notificator>
    );
  }
}
export default DirectoryDetails;
