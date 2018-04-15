import * as React from "react";
import { Query } from "react-apollo";

import { containerDetailsUrl } from ".";
import ContainerList from "../../components/ContainerList";
import Navigator from "../../components/Navigator";
import { getRootContainersQuery } from "./queries";

export const RootContainerList: React.StatelessComponent = () => (
  <Query query={getRootContainersQuery}>
    {({ data, loading, error }) => {
      if (error) {
        console.error(error);
        return;
      }
      return (
        <Navigator>
          {navigate => (
            <ContainerList
              containers={data.getRootContainers}
              loading={loading}
              onRowClick={id => event => navigate(containerDetailsUrl(id))}
            />
          )}
        </Navigator>
      );
    }}
  </Query>
);
export default RootContainerList;
