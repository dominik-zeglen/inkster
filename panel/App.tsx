import { InMemoryCache } from "apollo-cache-inmemory";
import { ApolloClient } from "apollo-client";
import { HttpLink } from "apollo-link-http";
import * as React from "react";
import { ApolloProvider } from "react-apollo";
import { ThemeProvider } from "react-jss";
import { BrowserRouter } from "react-router-dom";

import Navigator from "./components/Navigator";
import Layout from "./Layout";
import theme from "./theme";
import Router from "./views/Router";

const apolloClient = new ApolloClient({
  cache: new InMemoryCache(),
  link: new HttpLink({
    credentials: "same-origin",
    uri: "/panel/graphql"
  })
});

export const App: React.StatelessComponent = () => (
  <ApolloProvider client={apolloClient}>
    <BrowserRouter basename={"/panel/"}>
      <ThemeProvider theme={theme}>
        <Navigator>
          {navigate => (
            <Layout
              onContainersClick={() => navigate("/containers/")}
              onHomeClick={() => navigate("/")}
              onBrandClick={() => navigate("/")}
              onSearchSubmit={event => {
                event.preventDefault();
                const queryParam = new FormData(event.target).get("search");
                navigate(`/?search=${queryParam}`);
              }}
            >
              <Router />
            </Layout>
          )}
        </Navigator>
      </ThemeProvider>
    </BrowserRouter>
  </ApolloProvider>
);
export default App;
