import * as React from "react";
import { RouteComponentProps, withRouter } from "react-router";

interface NavigatorProps {
  children: ((
    navigate: (url: string, replace?: boolean) => void,
  ) => React.ReactElement);
}

export const Navigator = withRouter<NavigatorProps & RouteComponentProps<any>>(
  ({ children, history }) =>
    children((url: string, replace = false) =>
      replace ? history.replace(url) : history.push(url),
    ),
);
export default Navigator;
