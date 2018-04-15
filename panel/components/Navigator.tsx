/* Insolently stolen from mirumee/saleor
   Thank you patrys for effort that you
   have put in correcting my mistakes
   and teaching me that I should think 
   about few different approaches,
   before I start doing anything.
*/

import * as invariant from "invariant";
import * as PropTypes from "prop-types";
import * as React from "react";

interface NavigatorProps {
  children:
    | ((
        navigate: (url: string, replace?: boolean) => any
      ) => React.ReactElement<any>)
    | React.ReactNode;
}

export const Navigator: React.StatelessComponent<NavigatorProps> = (
  { children },
  { router }
) => {
  invariant(router, "You should not use <Navigator> outside a <Router>");
  const { history } = router;
  const navigate = (url, replace = false) =>
    replace ? history.replace(url) : history.push(url);

  if (typeof children === "function") {
    return children(navigate);
  }
  if (React.Children.count(children) > 0) {
    return React.Children.only(children);
  }
  return null;
};

Navigator.contextTypes = {
  router: PropTypes.shape({
    history: PropTypes.shape({
      push: PropTypes.func.isRequired,
      replace: PropTypes.func.isRequired
    }).isRequired
  })
};

export default Navigator;
