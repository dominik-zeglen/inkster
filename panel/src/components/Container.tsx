import * as React from "react";
import withStyles from "react-jss";

interface Props {
  width: "xs" | "sm" | "md" | "lg" | string;
}

const decorate = withStyles((theme: any) =>
  ["xs", "sm", "md", "lg", "xl"].reduce((prev, current) => {
    prev[current] = {
      [theme.breakpoints.up(current)]: {
        marginLeft: "auto",
        marginRight: "auto",
        maxWidth: theme.breakpoints.width(current)
      }
    };
    return prev;
  }, {})
);
export const Container =
  decorate <
  Props >
  (({ classes, width, ...props }) => (
    <div className={classes[width]} {...props} />
  ));
export default Container;
