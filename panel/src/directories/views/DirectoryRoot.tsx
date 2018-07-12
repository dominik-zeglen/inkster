import * as React from "react";
import { Mutation, Query } from "react-apollo";

import qRootDirectories from "../queries/qRootDirectories";
import mDirectoryCreate, {
  DirectoryCreateVariables
} from "../queries/mDirectoryCreate";
import DirectoryRootPage from "../components/DirectoryRootPage";
import Navigator from "../../components/Navigator";
import { urls } from "../../";

const dummy = () => {};

export const DirectoryRoot = () => (
  <Navigator>
    {navigate => {
      const handleRowClick = (id: string) => () =>
        navigate(urls.directoryDetails(id));
      const handleCreate = (data: { createContainer: { id: string } }) =>
        navigate(urls.directoryDetails(data.createContainer.id));
      return (
        <Query query={qRootDirectories} fetchPolicy="network-only">
          {({ data, error, loading }) => {
            if (error) {
              console.error(error);
              return <div>{JSON.stringify(error)}</div>;
            }
            return (
              <Mutation mutation={mDirectoryCreate} onCompleted={handleCreate}>
                {addDirectory => {
                  const handleAddDirectory = (
                    variables: DirectoryCreateVariables
                  ) => addDirectory({ variables });
                  return (
                    <DirectoryRootPage
                      directories={
                        data && data.getRootContainers
                          ? data.getRootContainers
                          : undefined
                      }
                      disabled={loading}
                      loading={loading}
                      onAdd={handleAddDirectory}
                      onNextPage={dummy}
                      onPreviousPage={dummy}
                      onRowClick={handleRowClick}
                    />
                  );
                }}
              </Mutation>
            );
          }}
        </Query>
      );
    }}
  </Navigator>
);
export default DirectoryRoot;
