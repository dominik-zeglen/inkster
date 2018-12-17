import * as React from "react";
import { Switch, Route, RouteComponentProps } from "react-router-dom";

import DirectoryDetailsComponent from "./DirectoryDetails";
import DirectoryRoot, { QueryParams } from "./DirectoryRoot";
import PageCreateComponent from "../../pages/views/PageCreate";
import { paths } from "../../urls";
import { parse as parseQs } from "qs";

const DirectoryList: React.StatelessComponent<RouteComponentProps<{}>> = ({
  location,
}) => {
  const qs = parseQs(location.search.substr(1));
  const params: QueryParams = {
    modal: qs.modal,
  };

  return <DirectoryRoot params={params} />;
};
const DirectoryDetails: React.StatelessComponent<RouteComponentProps<any>> = ({
  match,
}) => <DirectoryDetailsComponent id={decodeURIComponent(match.params.id)} />;
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
