import * as React from "react";
import { Route, Switch } from "react-router-dom";

import DirectorySection from "./directories/views";
import PageSection from './pages/views'

interface Props {
  match?: any;
}

export const App: React.StatelessComponent<Props> = () => {
  return (
    <Switch>
      <Route path={`/directories`} component={DirectorySection} />
      <Route path={`/pages`} component={PageSection} />
    </Switch>
  );
};
export default App;
