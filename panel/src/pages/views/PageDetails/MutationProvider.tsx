import * as React from "react";
import { MutationResult } from "react-apollo";
import { ApolloError } from "apollo-client";

import PageDeleteMutation from "../../queries/mPageDelete";
import PageUpdateMutation from "../../queries/mPageUpdate";
import {
  PageUpdateVariables,
  PageUpdate,
} from "../../queries/types/PageUpdate";
import {
  PageDelete,
  PageDeleteVariables,
} from "../../queries/types/PageDelete";

interface Props {
  children: ((
    props: {
      deletePage: {
        mutate: (variables: PageDeleteVariables) => void;
        opts: MutationResult<PageDelete>;
      };
      updatePage: {
        mutate: (variables: PageUpdateVariables) => void;
        opts: MutationResult<PageUpdate>;
      };
    },
  ) => React.ReactElement<any>);
  id: string;
  onError: (error: ApolloError) => void;
  onPageDelete: () => void;
  onPageUpdate: () => void;
}

export const MutationProvider: React.StatelessComponent<Props> = ({
  children,
  onPageDelete,
  onPageUpdate,
  onError,
}) => (
  <PageDeleteMutation onCompleted={onPageDelete} onError={onError}>
    {(deletePage, deletePageOpts) => {
      return (
        <PageUpdateMutation onCompleted={onPageUpdate} onError={onError}>
          {(updatePage, updatePageOpts) =>
            children({
              deletePage: {
                mutate: variables => deletePage({ variables }),
                opts: deletePageOpts,
              },
              updatePage: {
                mutate: variables => updatePage({ variables }),
                opts: updatePageOpts,
              },
            })
          }
        </PageUpdateMutation>
      );
    }}
  </PageDeleteMutation>
);
export default MutationProvider;
