import * as React from "react";

import Navigator from "./Navigator";

export const Link: React.StatelessComponent<
  React.HTMLProps<HTMLAnchorElement>
> = props => (
  <Navigator>
    {navigate => {
      const handleClick = (event: React.MouseEvent<any>) => {
        event.preventDefault();
        event.stopPropagation();
        if (props.href) {
          navigate(props.href);
        }
      };
      return <a {...props} onClick={handleClick} />;
    }}
  </Navigator>
);
export default Link;
