import * as React from "react";
import { Query } from "react-apollo";

import qRootDirectories from "../queries/qRootDirectories.gql";
import DirectoryRootPage from "../components/DirectoryRootPage";

export const DirectoryRoot: React.StatelessComponent = () => (
  <Query query={qRootDirectories}>
    {({ data, error, loading }) => {
      if (error) {
        console.error(error);
        return <div>{JSON.stringify(error)}</div>;
      }
      return (
        <DirectoryRootPage
          directories={
            data && data.getContainers ? data.getContainers : undefined
          }
          disabled={loading}
          loading={loading}
          onAdd={() => {}}
          onBack={() => window.history.back()}
          onNextPage={() => {}}
          onPreviousPage={() => {}}
          onRowClick={() => {}}
        />
      );
    }}
  </Query>
);
