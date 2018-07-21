import * as React from "react";
import { Switch, Route, RouteComponentProps } from "react-router-dom";

import DirectoryDetailsComponent from "./DirectoryDetails";
import DirectoryRoot from "./DirectoryRoot";
import PageCreateComponent from "../../pages/views/PageCreate";
import { unurlize } from "../../utils";

interface Props {
  match: any;
}

const DirectoryDetails: React.StatelessComponent<RouteComponentProps<any>> = ({
  match
}) => <DirectoryDetailsComponent id={unurlize(match.params.id)} />;
const PageCreate: React.StatelessComponent<RouteComponentProps<any>> = ({
  match
}) => <PageCreateComponent directory={unurlize(match.params.id)} />;

export const DirectorySection: React.StatelessComponent<Props> = ({
  match
}) => (
  <Switch>
    <Route exact={true} path={`${match.url}/`} component={DirectoryRoot} />
    <Route exact={true} path={`${match.url}/:id/`} component={DirectoryDetails} />
    <Route exact={true} path={`${match.url}/:id/createPage`} component={PageCreate} />
  </Switch>
);
export default DirectorySection;
