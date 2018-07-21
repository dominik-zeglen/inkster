import * as React from "react";
import { Query } from "react-apollo";
import { ApolloError } from "apollo-client";

import MutationProvider from "./MutationProvider";
import qDirectory from "../../queries/qDirectory";
import DirectoryDetailsPage from "../../components/DirectoryDetailsPage";
import Navigator from "../../../components/Navigator";
import { urls } from "../../../";
import { TransactionState } from "../../../";

const dummy = () => {};

interface Props {
  id: string;
}
interface State {
  transaction: TransactionState;
}
export class DirectoryDetails extends React.Component<Props, State> {
  state = {
    transaction: "default" as "default"
  };

  handleUpdate = () => {
    this.setState({ transaction: "success" });
    setTimeout(() => this.setState({ transaction: "default" }), 3000);
  };
  handleUpdateError = (event: ApolloError) => {
    this.setState({ transaction: "error" });
    setTimeout(() => this.setState({ transaction: "default" }), 3000);
  };

  render() {
    const { id } = this.props;
    return (
      <Navigator>
        {navigate => {
          const handleAddPage = () => navigate(urls.pageCreate(id))
          const handleRowClick = (id: string) => () =>
            navigate(urls.pageDetails(id));
          const handleDelete = () => navigate(urls.directoryDetails(), true)
          return (
            <Query
              query={qDirectory}
              variables={{ id }}
              fetchPolicy="network-only"
            >
              {({ data, error, loading }) => {
                if (error) {
                  console.error(error);
                  return <div>{JSON.stringify(error)}</div>;
                }
                return (
                  <MutationProvider
                    id={id}
                    onDirectoryUpdate={this.handleUpdate}
                    onDirectoryUpdateError={this.handleUpdateError}
                    onDirectoryDelete={handleDelete}
                    onDirectoryDeleteError={dummy}
                  >
                    {({ deleteDirectory, updateDirectory }) => {
                      return (
                        <DirectoryDetailsPage
                          directory={
                            data
                              ? data.getDirectory
                              : undefined
                          }
                          disabled={loading}
                          loading={loading}
                          transaction={
                            updateDirectory.loading
                              ? "loading"
                              : this.state.transaction
                          }
                          onAdd={handleAddPage}
                          onBack={
                            data && data.getDirectory
                              ? data.getDirectory.parent &&
                                data.getDirectory.parent.id
                                ? () =>
                                    navigate(
                                      urls.directoryDetails(
                                        data.getDirectory.parent.id
                                      )
                                    )
                                : () => navigate(urls.directoryDetails(""))
                              : () => window.history.back()
                          }
                          onDelete={deleteDirectory.mutate}
                          onNextPage={dummy}
                          onPreviousPage={dummy}
                          onRowClick={handleRowClick}
                          onSubmit={updateDirectory.mutate}
                        />
                      );
                    }}
                  </MutationProvider>
                );
              }}
            </Query>
          );
        }}
      </Navigator>
    );
  }
}
export default DirectoryDetails;
