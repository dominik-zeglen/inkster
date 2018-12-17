import * as React from "react";
import { MutationResult } from "react-apollo";

import DirectoryDeleteMutation from "../../queries/mDirectoryDelete";
import DirectoryUpdateMutation from "../../queries/mDirectoryUpdate";
import {
  DirectoryUpdateVariables,
  DirectoryUpdate,
} from "../../queries/types/DirectoryUpdate";
import { DirectoryDelete } from "../../queries/types/DirectoryDelete";
import PageCreateMutation from "../../queries/mPageCreate";
import {
  PageCreateVariables,
  PageCreate,
} from "../../queries/types/PageCreate";

interface Props {
  children: ((
    props: {
      createPage: {
        mutate: (variables: PageCreateVariables) => void;
        opts: MutationResult<PageCreate>;
      };
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
  onPageCreate: (data: PageCreate) => void;
}

export const MutationProvider: React.StatelessComponent<Props> = ({
  children,
  id,
  onDirectoryDelete,
  onDirectoryUpdate,
  onPageCreate,
}) => (
  <PageCreateMutation onCompleted={onPageCreate}>
    {(createPage, createPageOpts) => (
      <DirectoryDeleteMutation onCompleted={onDirectoryDelete}>
        {(deleteDirectory, deleteDirectoryOpts) => (
          <DirectoryUpdateMutation onCompleted={onDirectoryUpdate}>
            {(updateDirectory, updateDirectoryOpts) =>
              children({
                createPage: {
                  mutate: variables => createPage({ variables }),
                  opts: createPageOpts,
                },
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
    )}
  </PageCreateMutation>
);
export default MutationProvider;
