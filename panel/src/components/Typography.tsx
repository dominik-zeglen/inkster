import * as React from "react";
import * as classNames from "classnames";
import withStyles from "react-jss";

import { StandardProps } from "./";

export interface TypographyProps extends StandardProps {
  component?: string;
  variant?:
    | "anchor"
    | "body"
    | "caption"
    | "button"
    | "mainHeading"
    | "subHeading";
}

const decorate = withStyles(theme => ({
  anchor: {
    ...theme.typography.anchor
  },
  body: {
    ...theme.typography.body
  },
  caption: {
    ...theme.typography.caption
  },
  button: {
    ...theme.typography.button
  },
  mainHeading: {
    ...theme.typography.mainHeading
  },
  subHeading: {
    ...theme.typography.subHeading
  }
}));
export const Typography = decorate<TypographyProps>(
  ({ className, classes, component, variant = "body", ...props }) => {
    const Component = component || "div";
    return (
      <Component
        className={classNames([className, classes[variant]])}
        {...props}
      />
    );
  }
);
export default Typography;
