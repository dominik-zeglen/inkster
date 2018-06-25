import * as React from "react";
import { Query } from "react-apollo";
import gql from "graphql-tag";

import ContainerListPage from "../components/ContainerListPage";

const query = gql`
  query ChildContainerView($id: ID!) {
    getContainer(id: $id) {
      id
      name
      children {
        id
        name
      }
    }
  }
`;

interface Props {
  id?: string;
}

export const ContainerList: React.StatelessComponent<Props> = ({ id }) => (
  <Query query={query} variables={{ id }}>
    {({ data, error }) => {
      if (error) {
        return <span>{JSON.stringify(error)}</span>;
      }
      return (
        <ContainerListPage
          container={data && data.getContainer ? data.getContainer : undefined}
          containers={
            data && data.getContainer && data.getContainer.children ? data.getContainer.children : undefined
          }
          variant="child"
        />
      );
    }}
  </Query>
);
export default ContainerList;
