import * as React from "react";
import { MutationResult } from "react-apollo";

import DirectoryDeleteMutation from "../../queries/mDirectoryDelete";
import DirectoryUpdateMutation from "../../queries/mDirectoryUpdate";
import {
  DirectoryUpdateVariables,
  DirectoryUpdate,
} from "../../queries/types/DirectoryUpdate";
import { DirectoryDelete } from "../../queries/types/DirectoryDelete";

interface Props {
  children: ((
    props: {
      deleteDirectory: {
        mutate: () => void;
        opts: MutationResult<DirectoryDelete>;
      };
      updateDirectory: {
        mutate: (variables: DirectoryUpdateVariables) => void;
        opts: MutationResult<DirectoryUpdate>;
      };
    },
  ) => React.ReactElement<any>);
  id: string;
  onDirectoryDelete: () => void;
  onDirectoryUpdate: () => void;
}

export const MutationProvider: React.StatelessComponent<Props> = ({
  children,
  id,
  onDirectoryDelete,
  onDirectoryUpdate,
}) => (
  <DirectoryDeleteMutation onCompleted={onDirectoryDelete}>
    {(deleteDirectory, deleteDirectoryOpts) => (
      <DirectoryUpdateMutation onCompleted={onDirectoryUpdate}>
        {(updateDirectory, updateDirectoryOpts) =>
          children({
            deleteDirectory: {
              mutate: () => deleteDirectory({ variables: { id } }),
              opts: deleteDirectoryOpts,
            },
            updateDirectory: {
              mutate: variables =>
                updateDirectory({
                  variables,
                }),
              opts: updateDirectoryOpts,
            },
          })
        }
      </DirectoryUpdateMutation>
    )}
  </DirectoryDeleteMutation>
);
export default MutationProvider;
