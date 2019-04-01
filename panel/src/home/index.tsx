import * as React from "react";

import Navigator from "../components/Navigator";
import { ViewerQuery } from "./queries/viewer";
import HomePage from "./components/HomePage";
import urls from "../urls";
import { maybe } from "../utils";

export const Home: React.StatelessComponent = () => (
  <Navigator>
    {navigate => (
      <ViewerQuery>
        {viewer => (
          <HomePage
            disabled={viewer.loading}
            user={maybe(() => viewer.data.viewer)}
            onPageClick={id => navigate(urls.pageDetails(id))}
          />
        )}
      </ViewerQuery>
    )}
  </Navigator>
);
export default Home;
