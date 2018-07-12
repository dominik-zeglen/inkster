import * as React from "react";
import { Switch, Route, RouteComponentProps } from "react-router-dom";

import DirectoryDetailsComponent from "./DirectoryDetails";
import DirectoryRoot from "./DirectoryRoot";
import {unurlize} from '../../utils'

interface Props {
  match: any;
}

const DirectoryDetails: React.StatelessComponent<RouteComponentProps<any>> = ({
  match
}) => <DirectoryDetailsComponent id={unurlize(match.params.id)} />;

export const DirectorySection: React.StatelessComponent<Props> = ({
  match
}) => (
  <Switch>
    <Route exact={true} path={`${match.url}/`} component={DirectoryRoot} />
    <Route path={`${match.url}/:id/`} component={DirectoryDetails} />
  </Switch>
);
export default DirectorySection;
