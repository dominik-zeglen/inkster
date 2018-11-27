import { ApolloError } from "apollo-client";
import { DocumentNode } from "graphql";
import * as React from "react";
import {
  Mutation,
  MutationFn,
  MutationProps,
  MutationResult,
  MutationUpdaterFn,
} from "react-apollo";

import Notificator, { NotificationType } from "./components/Notificator";
import i18n from "./i18n";

interface TypedMutationInnerProps<TData, TVariables> {
  children: (
    mutateFn: MutationFn<TData, TVariables>,
    result: MutationResult<TData>,
  ) => React.ReactNode;
  onCompleted?: (data: TData) => void;
  onError?: (error: ApolloError) => void;
  variables?: TVariables;
}

export function TypedMutation<TData, TVariables>(
  mutation: DocumentNode,
  update?: MutationUpdaterFn<TData>,
) {
  const StrictTypedMutation: React.ComponentType<
    MutationProps<TData, TVariables>
  > = Mutation;
  return ({
    children,
    onCompleted,
    onError,
    variables,
  }: TypedMutationInnerProps<TData, TVariables>) => (
    <Notificator>
      {notify => (
        <StrictTypedMutation
          mutation={mutation}
          onCompleted={onCompleted}
          onError={err => {
            const msg = i18n.t("Something went wrong: {{ message }}", {
              message: err.message,
            });
            notify({ text: msg, type: NotificationType.ERROR });
            if (onError) {
              onError(err);
            }
          }}
          variables={variables}
          update={update}
        >
          {children}
        </StrictTypedMutation>
      )}
    </Notificator>
  );
}
