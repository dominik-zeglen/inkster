import { ApolloClient } from "apollo-client";
import { ApolloProvider } from "react-apollo";
import * as React from "react";
import { BrowserRouter } from "react-router-dom";
import { HttpLink } from "apollo-link-http";
import { InMemoryCache } from "apollo-cache-inmemory";
import { render } from "react-dom";
import { ThemeProvider } from "react-jss";

import App from "./App";
import AppRoot from "./AppRoot";
import GlobalStylesheet from "./Stylesheet";
import theme from "./theme";
import UploadProvider from "./UploadProvider";
import LoaderOverlay from "./components/LoaderOverlay";
import {DateProvider} from "./components/Date";
import { urlize } from "./utils";

const apolloClient = new ApolloClient({
  cache: new InMemoryCache(),
  link: new HttpLink({
    credentials: "same-origin",
    uri: "/graphql/"
  })
});

render(
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
                <AppRoot>
                  <App />
                </AppRoot>
                {uploadState.active && (
                  <LoaderOverlay progress={uploadState.progress} />
                )}
              </>
            </ThemeProvider>
          </BrowserRouter>
        </ApolloProvider>
      )}
    </UploadProvider>
  </DateProvider>,
  document.querySelector("#root")
);

export const urls = {
  directoryDetails: (id?: string) => `/directories/${id ? urlize(id) : ""}`,
  pageCreate: (id: string) => `/directories/${id}/createPage`,
  pageDetails: (id: string) => `/pages/${urlize(id)}`
};

export type TransactionState = "default" | "loading" | "success" | "error";
export interface PaginationInfo {
  hasNextPage: boolean;
  hasPreviousPage: boolean;
}
export interface PaginatedListProps {
  pageInfo?: PaginationInfo;
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
