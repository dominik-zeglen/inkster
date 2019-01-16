import * as React from "react";
import { Route, Switch } from "react-router-dom";

import DirectorySection from "./directories/views";
import PageSection from "./pages/views";
import UserSection from "./users/views";
import WebsiteSection from "./website/views";
import Home from "./home";
import {
  directorySection,
  home,
  pageSection,
  websiteSettingsSection,
  userSection,
} from "./urls";

interface Props {
  match?: any;
}

export const App: React.StatelessComponent<Props> = () => {
  return (
    <Switch>
      <Route path={home} exact={true} component={Home} />
      <Route path={directorySection} component={DirectorySection} />
      <Route path={pageSection} component={PageSection} />
      <Route path={userSection} component={UserSection} />
      <Route path={websiteSettingsSection} component={WebsiteSection} />
    </Switch>
  );
};
export default App;
