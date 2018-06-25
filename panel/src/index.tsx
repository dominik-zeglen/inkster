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
