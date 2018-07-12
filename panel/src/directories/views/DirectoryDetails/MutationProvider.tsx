import * as React from "react";
import { Mutation } from "react-apollo";
import { ApolloError } from "apollo-client";

import mDirectoryDelete from "../../queries/mDirectoryDelete";
import mDirectoryUpdate, {
  variables as updateDirectoryVariables
} from "../../queries/mDirectoryUpdate";

interface Props {
  children:
    | ((
        props: {
          deleteDirectory: {
            mutate: () => void;
            loading: boolean;
          };
          updateDirectory: {
            mutate: (
              variables: Exclude<updateDirectoryVariables, { id: string }>
            ) => void;
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
  onDirectoryDelete: () => void;
  onDirectoryDeleteError: (error: ApolloError) => void;
  onDirectoryUpdate: () => void;
  onDirectoryUpdateError: (error: ApolloError) => void;
}

export const MutationProvider: React.StatelessComponent<Props> = ({
  children,
  id,
  onDirectoryDelete,
  onDirectoryDeleteError,
  onDirectoryUpdate,
  onDirectoryUpdateError
}) => (
  <Mutation mutation={mDirectoryDelete} onCompleted={onDirectoryDelete}>
    {(
      deleteDirectory,
      { called, data, error, loading: deleteDirectoryLoading }
    ) => {
      if (error) {
        onDirectoryDeleteError(error);
      }
      return (
        <Mutation mutation={mDirectoryUpdate} onCompleted={onDirectoryUpdate}>
          {(
            updateDirectory,
            { called, data, error, loading: updateDirectoryLoading }
          ) => {
            if (error) {
              onDirectoryUpdateError(error);
            }
            return children && typeof children === "function"
              ? children({
                  deleteDirectory: {
                    mutate: () => deleteDirectory({ variables: { id } }),
                    loading: deleteDirectoryLoading
                  },
                  updateDirectory: {
                    mutate: (
                      variables: Exclude<
                        updateDirectoryVariables,
                        { id: string }
                      >
                    ) =>
                      updateDirectory({
                        variables: { id, ...(variables as any) }
                      }),
                    loading: updateDirectoryLoading
                  },
                  formErrors:
                    data && data.updateContainer && data.updateContainer.errors
                      ? data.updateContainers.errors
                      : undefined
                })
              : null;
          }}
        </Mutation>
      );
    }}
  </Mutation>
);
export default MutationProvider;
