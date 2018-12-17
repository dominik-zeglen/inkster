import * as React from "react";

import Navigator from "../components/Navigator";
import { ViewerQuery } from "./queries/viewer";
import HomePage from "./components/HomePage";
import urls from "../urls";

export const Home: React.StatelessComponent = () => (
  <Navigator>
    {navigate => (
      <ViewerQuery>
        {viewer => (
          <HomePage
            disabled={viewer.loading}
            user={viewer.data.viewer}
            onPageClick={id => navigate(urls.pageDetails(id))}
          />
        )}
      </ViewerQuery>
    )}
  </Navigator>
);
export default Home;
