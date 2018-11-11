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
      { error: deleteDirectoryError, loading: deleteDirectoryLoading }
    ) => {
      if (deleteDirectoryError) {
        onDirectoryDeleteError(deleteDirectoryError);
      }
      return (
        <Mutation mutation={mDirectoryUpdate} onCompleted={onDirectoryUpdate}>
          {(
            updateDirectory,
            { data, error: updateDirectoryError, loading: updateDirectoryLoading }
          ) => {
            if (updateDirectoryError) {
              onDirectoryUpdateError(updateDirectoryError);
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
                    data && data.updateDirectory && data.updateDirectory.errors
                      ? data.updateDirectory.errors
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
