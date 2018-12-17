import * as React from "react";
import { Switch, Route, RouteComponentProps } from "react-router-dom";

import PageDetailsComponent from "./PageDetails";
import { paths } from "../../urls";

const PageDetails: React.StatelessComponent<RouteComponentProps<any>> = ({
  match,
}) => <PageDetailsComponent id={decodeURIComponent(match.params.id)} />;

export const PageSection: React.StatelessComponent = () => (
  <Switch>
    <Route path={paths.pageDetails(":id")} component={PageDetails} />
  </Switch>
);
export default PageSection;
