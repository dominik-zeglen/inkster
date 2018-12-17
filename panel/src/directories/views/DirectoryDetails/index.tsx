import * as React from "react";

import MutationProvider from "./MutationProvider";
import Directory from "../../queries/qDirectory";
import DirectoryDetailsPage from "../../components/DirectoryDetailsPage";
import Navigator from "../../../components/Navigator";
import Notificator from "../../../components/Notificator";
import urls from "../../../urls";
import i18n from "../../../i18n";
import { TransactionState } from "../../../";
import { maybe, mergeQs } from "../../../utils";
import { Modal } from "../../../types";
import ActionDialog from "../../../components/ActionDialog";

const dummy = () => {};

export type QueryParams = Partial<Modal<"remove">>;
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
              const handleAddPage = () => navigate(urls.pageCreate(id));
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
              return (
                <Directory variables={{ id }}>
                  {directory => (
                    <MutationProvider
                      id={id}
                      onDirectoryUpdate={this.handleUpdate}
                      onDirectoryDelete={handleDelete}
                    >
                      {({ deleteDirectory, updateDirectory }) => (
                        <>
                          <DirectoryDetailsPage
                            directory={maybe(() => directory.data.getDirectory)}
                            disabled={directory.loading}
                            loading={directory.loading}
                            transaction={
                              updateDirectory.opts.loading
                                ? "loading"
                                : this.state.transaction
                            }
                            onAdd={handleAddPage}
                            onBack={
                              directory.data && directory.data.getDirectory
                                ? directory.data.getDirectory.parent &&
                                  directory.data.getDirectory.parent.id
                                  ? () =>
                                      navigate(
                                        urls.directoryDetails(
                                          directory.data.getDirectory.parent.id,
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
                            onNextPage={dummy}
                            onPreviousPage={dummy}
                            onRowClick={handleRowClick}
                            onSubmit={formData =>
                              updateDirectory.mutate({
                                ...formData,
                                id,
                              })
                            }
                          />
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
