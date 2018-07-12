import * as React from "react";
import { RouteComponentProps, withRouter } from "react-router";

interface NavigatorProps {
  children:
    | ((
        navigate: (url: string, replace?: boolean) => any
      ) => React.ReactElement<any>)
    | React.ReactNode;
}

export const Navigator = withRouter<NavigatorProps & RouteComponentProps<any>>(
  ({ children, history }) => {
    const navigate = (url: string, replace = false) =>
      replace ? history.replace(url) : history.push(url);

    if (typeof children === "function") {
      return children(navigate);
    }
    if (React.Children.count(children) > 0) {
      return React.Children.only(children);
    }
    return null;
  }
);
export default Navigator;
