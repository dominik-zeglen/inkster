import * as React from "react";
import { Switch, Route, RouteComponentProps } from "react-router-dom";

import DirectoryDetailsComponent, {
  QueryParams as DirectoryDetailsQueryParams,
} from "./DirectoryDetails";
import DirectoryRoot, {
  QueryParams as DirectoryRootQueryParams,
} from "./DirectoryRoot";
import PageCreateComponent from "../../pages/views/PageCreate";
import { paths } from "../../urls";
import { parse as parseQs } from "qs";

const DirectoryList: React.StatelessComponent<RouteComponentProps<{}>> = ({
  location,
}) => {
  const qs = parseQs(location.search.substr(1));
  const params: DirectoryRootQueryParams = {
    modal: qs.modal,
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
  };

  return (
    <DirectoryDetailsComponent
      id={decodeURIComponent(match.params.id)}
      params={params}
    />
  );
};
const PageCreate: React.StatelessComponent<RouteComponentProps<any>> = ({
  match,
}) => <PageCreateComponent directory={decodeURIComponent(match.params.id)} />;

export const DirectorySection: React.StatelessComponent = () => (
  <Switch>
    <Route exact={true} path={paths.directoryList} component={DirectoryList} />
    <Route
      exact={true}
      path={paths.directoryDetails(":id")}
      component={DirectoryDetails}
    />
    <Route exact={true} path={paths.pageCreate(":id")} component={PageCreate} />
  </Switch>
);
export default DirectorySection;
