import { ApolloError } from "apollo-client";
import { DocumentNode } from "graphql";
import * as React from "react";
import {
  Mutation,
  MutationFn,
  MutationResult,
  MutationUpdaterFn,
  Query,
  QueryResult,
} from "react-apollo";
import gql from "graphql-tag";

import Notificator, { NotificationType } from "./components/Notificator";
import i18n from "./i18n";
import AppProgress from "./components/AppProgress";

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
  class StrictTypedMutation extends Mutation<TData, TVariables> {}
  return (props: TypedMutationInnerProps<TData, TVariables>) => {
    const {
      children,
      onCompleted,
      onError,
      variables,
    } = props as JSX.LibraryManagedAttributes<
      typeof StrictTypedMutation,
      typeof props
    >;
    return (
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
  };
}

interface TypedQueryInnerProps<TData, TVariables> {
  children: (result: QueryResult<TData, TVariables>) => React.ReactNode;
  displayLoader?: boolean;
  skip?: boolean;
  variables?: TVariables;
}

interface QueryProgressProps {
  loading: boolean;
  onLoading: () => void;
  onCompleted: () => void;
}

class QueryProgress extends React.Component<QueryProgressProps, {}> {
  componentDidMount() {
    const { loading, onLoading } = this.props;
    if (loading) {
      onLoading();
    }
  }

  componentDidUpdate(prevProps) {
    const { loading, onLoading, onCompleted } = this.props;
    if (prevProps.loading !== loading) {
      if (loading) {
        onLoading();
      } else {
        onCompleted();
      }
    }
  }

  render() {
    return this.props.children;
  }
}

export function TypedQuery<TData, TVariables>(query: DocumentNode) {
  class StrictTypedQuery extends Query<TData, TVariables> {}
  return (props: TypedQueryInnerProps<TData, TVariables>) => {
    const {
      children,
      displayLoader,
      skip,
      variables,
    } = props as JSX.LibraryManagedAttributes<
      typeof StrictTypedQuery,
      typeof props
    >;
    return (
      <AppProgress>
        {({ funcs: changeProgressState }) => (
          <Notificator>
            {notify => (
              <StrictTypedQuery
                fetchPolicy="cache-and-network"
                query={query}
                variables={variables}
                skip={skip}
              >
                {queryData => {
                  if (queryData.error) {
                    const msg = i18n.t("Something went wrong: {{ message }}", {
                      message: queryData.error.message,
                    });
                    notify({ text: msg, type: NotificationType.ERROR });
                  }

                  if (displayLoader) {
                    return (
                      <QueryProgress
                        loading={queryData.loading}
                        onCompleted={changeProgressState.disable}
                        onLoading={changeProgressState.enable}
                      >
                        {children(queryData)}
                      </QueryProgress>
                    );
                  }

                  return children(queryData);
                }}
              </StrictTypedQuery>
            )}
          </Notificator>
        )}
      </AppProgress>
    );
  };
}

export const pageInfoFragment = gql`
  fragment PageInfoFragment on PageInfo {
    startCursor
    endCursor
    hasPreviousPage
    hasNextPage
  }
`;
