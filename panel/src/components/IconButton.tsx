import * as React from "react";
import withStyles from "react-jss";

import { StandardProps } from "./";

interface Props extends StandardProps {
  disabled?: boolean;
  icon: React.StatelessComponent<{ color?: string } & StandardProps>;
}

const decorate = withStyles((theme: any) => ({
  root: {
    borderRadius: "100%",
    padding: theme.spacing / 2
  },
  container: {
    display: "flex",
    alignItems: "center",
    width: theme.spacing * 2,
    height: theme.spacing * 2
  },
  icon: {
    color: theme.colors.black,
    cursor: "pointer",
    transitionDuration: theme.transition.time,
    "&:hover, &:focus": {
      color: theme.colors.secondary.main
    }
  },
  disabledIcon: {
    color: theme.colors.disabled,
    cursor: "pointer"
  }
}));
export const IconButton =
  decorate <
  Props >
  (({ classes, disabled, icon, onClick, ...props }) => {
    const Icon = icon;
    const shouldBeDisabled = disabled || !onClick;
    return (
      <div
        className={classes.root}
        onClick={shouldBeDisabled ? undefined : onClick}
        {...props}
      >
        <div className={classes.container}>
          <Icon
            className={shouldBeDisabled ? classes.disabledIcon : classes.icon}
          />
        </div>
      </div>
    );
  });
export default IconButton;
