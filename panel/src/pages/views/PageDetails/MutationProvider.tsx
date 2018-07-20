import * as React from "react";
import { Mutation } from "react-apollo";
import { ApolloError } from "apollo-client";

import mPageDelete from "../../queries/mPageDelete";
import mPageUpdate, {
  variables as updatePageVariables
} from "../../queries/mPageUpdate";

interface Props {
  children:
    | ((
        props: {
          deletePage: {
            mutate: () => void;
            loading: boolean;
          };
          updatePage: {
            mutate: (variables: updatePageVariables) => void;
            loading: boolean;
          };
          formErrors: Array<{
            field: string;
            message: string;
          }>;
        }
      ) => React.ReactElement<any>)
    | React.ReactNode;
  id: string;
  onError: (error: ApolloError) => void;
  onPageDelete: () => void;
  onPageUpdate: () => void;
}

export const MutationProvider: React.StatelessComponent<Props> = ({
  children,
  id,
  onPageDelete,
  onPageUpdate,
  onError
}) => (
  <Mutation mutation={mPageDelete} onCompleted={onPageDelete} onError={onError}>
    {(deletePage, { data, error, loading: deletePageLoading }) => {
      return (
        <Mutation
          mutation={mPageUpdate}
          onCompleted={onPageUpdate}
          onError={onError}
        >
          {(updatePage, { data, error, loading: updatePageLoading }) => {
            return children && typeof children === "function"
              ? children({
                  deletePage: {
                    mutate: () => deletePage({ variables: { id } }),
                    loading: deletePageLoading
                  },
                  formErrors: [],
                  updatePage: {
                    mutate: (variables: updatePageVariables) => updatePage({variables}),
                    loading: updatePageLoading
                  }
                })
              : null;
          }}
        </Mutation>
      );
    }}
  </Mutation>
);
export default MutationProvider;
