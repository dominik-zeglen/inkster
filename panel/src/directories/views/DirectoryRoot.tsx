import * as React from "react";
import { Mutation, Query } from "react-apollo";

import qRootDirectories from "../queries/qRootDirectories";
import mDirectoryCreate, {
  DirectoryCreateVariables,
} from "../queries/mDirectoryCreate";
import DirectoryRootPage from "../components/DirectoryRootPage";
import Navigator from "../../components/Navigator";
import Notificator from "../../components/Notificator";
import urls from "../../urls";
import i18n from "../../i18n";

const dummy = () => {};

export const DirectoryRoot = () => (
  <Notificator>
    {notify => (
      <Navigator>
        {navigate => {
          const handleRowClick = (id: string) => () =>
            navigate(urls.directoryDetails(id));
          const handleCreate = (data: { createDirectory: { id: string } }) => {
            notify({
              text: i18n.t("Created directory", {
                context: "notification",
              }),
            });
            navigate(urls.directoryDetails(data.createDirectory.id));
          };
          return (
            <Query query={qRootDirectories} fetchPolicy="network-only">
              {({ data, error, loading }) => {
                if (error) {
                  console.error(error);
                  return <div>{JSON.stringify(error)}</div>;
                }
                return (
                  <Mutation
                    mutation={mDirectoryCreate}
                    onCompleted={handleCreate}
                  >
                    {addDirectory => {
                      const handleAddDirectory = (
                        variables: DirectoryCreateVariables,
                      ) => addDirectory({ variables });
                      return (
                        <DirectoryRootPage
                          directories={
                            data ? data.getRootDirectories : undefined
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
    )}
  </Notificator>
);
export default DirectoryRoot;
