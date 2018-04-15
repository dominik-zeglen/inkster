import * as React from "react";
import { Mutation, Query } from "react-apollo";

import { containerDetailsUrl, containerEditUrl } from ".";
import ContainerList from "../../components/ContainerList";
import ContainerProperties from "../../components/ContainerProperties";
import Navigator from "../../components/Navigator";
import {
  getContainerChildrenQuery,
  getContainerDetailsQuery,
  removeContainerMutation
} from "./queries";

interface ContainerSectionProps {
  match: any;
}

export const ContainerDetails: React.StatelessComponent<
  ContainerSectionProps
> = ({ match }) => {
  const id = parseInt(match.params.id, 0);

  return (
    <Navigator>
      {navigate => (
        <>
          <Query query={getContainerDetailsQuery} variables={{ id }}>
            {({ data, loading, error }) => {
              if (error) {
                console.error(error);
                return;
              }
              return (
                <Mutation mutation={removeContainerMutation} variables={{ id }}>
                  {(
                    handleRemove,
                    { called: calledRemove, loading: loadingRemove, error }
                  ) => {
                    if (error) {
                      return;
                    }
                    return (
                      <ContainerProperties
                        container={loading ? null : data.getContainer}
                        loading={loading}
                        onEdit={() => navigate(containerEditUrl(id))}
                        onRemove={handleRemove}
                      />
                    );
                  }}
                </Mutation>
              );
            }}
          </Query>
          <Query query={getContainerChildrenQuery} variables={{ id }}>
            {({ data, loading, error }) => {
              if (error) {
                console.error(error);
                return;
              }
              return (
                <ContainerList
                  containers={loading ? [] : data.getContainer.children}
                  loading={loading}
                  onRowClick={id => event => navigate(containerDetailsUrl(id))}
                />
              );
            }}
          </Query>
        </>
      )}
    </Navigator>
  );
};
export default ContainerDetails;
