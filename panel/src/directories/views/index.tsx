import * as React from "react";
import { Switch, Route, RouteComponentProps } from "react-router-dom";

import DirectoryDetailsComponent, {
  QueryParams as DirectoryDetailsQueryParams,
} from "./DirectoryDetails";
import DirectoryRoot, {
  QueryParams as DirectoryRootQueryParams,
} from "./DirectoryRoot";
import { paths } from "../../urls";
import { parse as parseQs } from "qs";

const DirectoryList: React.StatelessComponent<RouteComponentProps<{}>> = ({
  location,
}) => {
  const qs = parseQs(location.search.substr(1));
  const params: DirectoryRootQueryParams = {
    modal: qs.modal,
    after: qs.after,
    before: qs.before,
  };

  return <DirectoryRoot params={params} />;
};
const DirectoryDetails: React.StatelessComponent<RouteComponentProps<any>> = ({
  location,
  match,
}) => {
  const qs = parseQs(location.search.substr(1));
  const params: DirectoryDetailsQueryParams = {
    modal: qs.modal,
    after: qs.after,
    before: qs.before,
  };

  return (
    <DirectoryDetailsComponent
      id={decodeURIComponent(match.params.id)}
      params={params}
    />
  );
};

export const DirectorySection: React.StatelessComponent = () => (
  <Switch>
    <Route exact={true} path={paths.directoryList} component={DirectoryList} />
    <Route
      exact={true}
      path={paths.directoryDetails(":id")}
      component={DirectoryDetails}
    />
  </Switch>
);
export default DirectorySection;
