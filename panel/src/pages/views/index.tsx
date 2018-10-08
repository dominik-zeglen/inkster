import * as React from "react";
import { Switch, Route, RouteComponentProps } from "react-router-dom";

import PageDetailsComponent from "./PageDetails";

interface Props {
  match: any;
}

const PageDetails: React.StatelessComponent<RouteComponentProps<any>> = ({
  match
}) => <PageDetailsComponent id={decodeURIComponent(match.params.id)} />;

export const PageSection: React.StatelessComponent<Props> = ({ match }) => (
  <Switch>
    <Route path={`${match.url}/:id/`} component={PageDetails} />
  </Switch>
);
export default PageSection;
