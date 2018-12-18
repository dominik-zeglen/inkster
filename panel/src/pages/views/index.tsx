import * as React from "react";
import { Switch, Route, RouteComponentProps } from "react-router-dom";
import { parse as parseQs } from "qs";

import PageDetailsComponent, {
  QueryParams as PageDetailsQueryParams,
} from "./PageDetails";
import { paths } from "../../urls";

const PageDetails: React.StatelessComponent<RouteComponentProps<any>> = ({
  match,
  location,
}) => {
  const qs = parseQs(location.search.substr(1));
  const params: PageDetailsQueryParams = {
    modal: qs.modal,
  };
  return (
    <PageDetailsComponent
      id={decodeURIComponent(match.params.id)}
      params={params}
    />
  );
};

export const PageSection: React.StatelessComponent = () => (
  <Switch>
    <Route path={paths.pageDetails(":id")} component={PageDetails} />
  </Switch>
);
export default PageSection;
