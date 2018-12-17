import * as React from "react";

import MutationProvider from "./MutationProvider";
import Directory from "../../queries/qDirectory";
import DirectoryDetailsPage from "../../components/DirectoryDetailsPage";
import Navigator from "../../../components/Navigator";
import Notificator from "../../../components/Notificator";
import urls from "../../../urls";
import i18n from "../../../i18n";
import { TransactionState } from "../../../";
import { maybe } from "../../../utils";

const dummy = () => {};

interface Props {
  id: string;
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
    const { id } = this.props;
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
                          onDelete={deleteDirectory.mutate}
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
