import { ApolloClient } from "apollo-client";
import { ApolloProvider } from "react-apollo";
import * as React from "react";
import { BrowserRouter } from "react-router-dom";
import { HttpLink } from "apollo-link-http";
import { InMemoryCache } from "apollo-cache-inmemory";
import { render } from "react-dom";
import { ThemeProvider } from "react-jss";

import App from "./App";
import GlobalStylesheet from "./Stylesheet";
import theme from "./theme";

const apolloClient = new ApolloClient({
  cache: new InMemoryCache(),
  link: new HttpLink({
    credentials: "same-origin",
    uri: "/graphql/"
  })
});

render(
  <ApolloProvider client={apolloClient}>
    <BrowserRouter basename="/panel/">
      <ThemeProvider theme={theme}>
        <>
          <GlobalStylesheet />
          <App />
        </>
      </ThemeProvider>
    </BrowserRouter>
  </ApolloProvider>,
  document.querySelector("#root")
);

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
}
export interface ListViewProps extends ViewProps, PaginatedListProps {
  onAdd: () => void;
}
export interface FormViewProps extends ViewProps {
  transaction: TransactionState;
  onBack: () => void;
  onDelete: () => void;
  onSubmit: () => void;
}
