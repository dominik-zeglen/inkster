import * as React from "react";
import { Route, Switch } from "react-router-dom";

import AppRoot from './components/AppRoot'
import ContainerSection from './containers/views'

export const App: React.StatelessComponent<{}> = () => (
  <AppRoot>
    <Switch>
      <Route path="/container" component={ContainerSection} />
    </Switch>
  </AppRoot>
);
export default App;
