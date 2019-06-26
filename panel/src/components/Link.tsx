import createUseStyles, { css } from "aurora-ui-kit/dist/utils/jss";
import * as React from "react";
import classNames from "classnames";

import Navigator from "./Navigator";

const useStyles = createUseStyles({
  root: css`
    color: unset;
  `,
});

export const Link: React.StatelessComponent<
  React.HTMLProps<HTMLAnchorElement>
> = ({ className, ...props }) => {
  const classes = useStyles();

  return (
    <Navigator>
      {navigate => {
        const handleClick = (event: React.MouseEvent<any>) => {
          event.preventDefault();
          event.stopPropagation();
          if (props.href) {
            navigate(props.href);
          }
        };
        return (
          <a
            className={classNames(classes.root, className)}
            {...props}
            onClick={handleClick}
          />
        );
      }}
    </Navigator>
  );
};
export default Link;
