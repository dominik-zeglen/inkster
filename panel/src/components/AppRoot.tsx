import * as React from "react";
import withStyles from "react-jss";

interface Props {}

const decorate = withStyles((theme: any) => ({}));
export const AppRoot =
  decorate < Props > (({ children }) => <div>{children}</div>);
export default AppRoot;
