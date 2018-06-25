import * as React from "react";
import { Query } from "react-apollo";
import gql from "graphql-tag";

import ContainerListPage from "../components/ContainerListPage";

const query = gql`
  query RootContainersView {
    getRootContainers {
      id
      name
    }
  }
`;

export const RootContainerList: React.StatelessComponent = () => (
  <Query query={query}>
    {({ data, error }) => {
      if (error) {
        return <span>{JSON.stringify(error)}</span>;
      }
      return (
        <ContainerListPage
          containers={
            data && data.getRootContainers ? data.getRootContainers : undefined
          }
          variant="root"
        />
      );
    }}
  </Query>
);
export default RootContainerList;
