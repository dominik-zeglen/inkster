import * as React from "react";
import { Route, Switch } from "react-router";

import ContainerDetails from "./ContainerDetails";
import RootContainerList from "./RootContainerList";

export const ContainerSection: React.StatelessComponent = () => (
  <Switch>
    <Route exact path="/containers/:id" component={ContainerDetails} />
    <Route exact path="/containers/" component={RootContainerList} />
  </Switch>
);

export const containerDetailsUrl = (id: number) => `/containers/${id}/`;
export const containerEditUrl = (id: number) => `/containers/${id}/edit/`;

export default ContainerSection;
