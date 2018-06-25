import * as React from "react";
import { Route, Switch } from "react-router-dom";

import RootContainerList from "./RootContainerList";
import ContainerList from "./ContainerList";

const AugmentedContainerList: React.StatelessComponent<{ match: any }> = ({
  match
}) => <ContainerList id={match.params.id} />;

export const CategorySection: React.StatelessComponent<{ match: any }> = (
  props,
  { match }
) => (
  <Switch>
    <Route exact={true} path="/container/" component={RootContainerList} />
    <Route
      exact={true}
      path="/container/:id/"
      component={AugmentedContainerList}
    />
  </Switch>
);
export default CategorySection;
