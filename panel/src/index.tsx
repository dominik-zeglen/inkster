import { ApolloClient } from "apollo-client";
import { setContext } from "apollo-link-context";
import { ErrorResponse, onError } from "apollo-link-error";
import { ApolloProvider } from "react-apollo";
import * as React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import { HttpLink } from "apollo-link-http";
import { InMemoryCache, defaultDataIdFromObject } from "apollo-cache-inmemory";
import { render } from "react-dom";
import { ThemeProvider } from "react-jss";

import App from "./App";
import AppRoot from "./AppRoot";
import GlobalStylesheet from "./Stylesheet";
import theme from "./theme";
import UploadProvider from "./UploadProvider";
import LoaderOverlay from "./components/LoaderOverlay";
import { DateProvider } from "./components/Date";
import AuthProvider, {
  getAuthToken,
  removeAuthToken,
} from "./auth/components/AuthProvider";
import Login from "./auth/views/Login";
import PasswordRecovery from "./auth/views/PasswordRecovery";
import { NotificationProvider } from "./components/Notificator";
import urls from "./urls";
import { AppProgressProvider } from "./components/AppProgress";

interface ResponseError extends ErrorResponse {
  networkError?: Error & {
    statusCode?: number;
    bodyText?: string;
  };
}

const invalidTokenLink = onError((error: ResponseError) => {
  if (error.networkError && error.networkError.statusCode === 401) {
    removeAuthToken();
  }
});

const authLink = setContext((_, context) => {
  const authToken = getAuthToken();
  return {
    ...context,
    headers: {
      ...context.headers,
      Authorization: authToken ? `Bearer ${authToken}` : null,
    },
  };
});

const apolloClient = new ApolloClient({
  cache: new InMemoryCache({
    dataIdFromObject: (obj: any) => {
      if (["Website"].indexOf(obj.__typename) !== -1) {
        return obj.__typename;
      }
      return defaultDataIdFromObject(obj);
    },
  }),
  link: invalidTokenLink.concat(
    authLink.concat(
      new HttpLink({
        credentials: "same-origin",
        uri: "/graphql/",
      }),
    ),
  ),
});

render(
  <AppProgressProvider>
    <DateProvider>
      <UploadProvider>
        {uploadState => (
          <ApolloProvider client={apolloClient}>
            <BrowserRouter
              basename={process.env.NODE_ENV === "production" ? "/panel/" : "/"}
            >
              <ThemeProvider theme={theme}>
                <>
                  <GlobalStylesheet />
                  <NotificationProvider>
                    <AuthProvider>
                      {({
                        hasToken,
                        isAuthenticated,
                        loginLoading,
                        tokenVerifyLoading,
                      }) =>
                        isAuthenticated ? (
                          <>
                            <AppRoot>
                              <App />
                            </AppRoot>
                            {uploadState.active && (
                              <LoaderOverlay progress={uploadState.progress} />
                            )}
                          </>
                        ) : hasToken && tokenVerifyLoading ? (
                          <span />
                        ) : (
                          <Switch>
                            <Route
                              path={urls.passwordRecovery}
                              component={PasswordRecovery}
                            />
                            <Route
                              component={() => <Login loading={loginLoading} />}
                            />
                          </Switch>
                        )
                      }
                    </AuthProvider>
                  </NotificationProvider>
                </>
              </ThemeProvider>
            </BrowserRouter>
          </ApolloProvider>
        )}
      </UploadProvider>
    </DateProvider>
  </AppProgressProvider>,
  document.querySelector("#root"),
);

export type TransactionState = "default" | "loading" | "success" | "error";
export interface PaginationInfo {
  hasNextPage: boolean;
  hasPreviousPage: boolean;
}
export interface PaginatedListProps {
  pageInfo: PaginationInfo;
  onNextPage: () => void;
  onPreviousPage: () => void;
  onRowClick: (id: string) => () => void;
}
export interface ViewProps {
  disabled: boolean;
  loading: boolean;
  title?: string;
}
export interface ListViewProps<T> extends ViewProps, PaginatedListProps {
  onAdd: (data?: T) => void;
}
export interface FormViewProps<T> extends ViewProps {
  transaction: TransactionState;
  onBack: () => void;
  onDelete?: () => void;
  onSubmit: (data: T) => void;
}

export const PAGINATE_BY = 10;
