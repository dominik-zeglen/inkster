import * as React from "react";
import withStyles from "react-jss";

import { StandardProps } from "./";

const decorate = withStyles(
  (theme: any) => ({
    "@keyframes skeleton-animation": {
      "0%": {
        opacity: 0.6
      },
      "100%": {
        opacity: 1
      }
    },
    skeleton: {
      animation: "skeleton-animation .75s linear infinite forwards alternate",
      background: '#f2f2f2',
      borderRadius: 4,
      display: "block",
      height: "0.8em",
      margin: "0.2em 0",
      width: "100%"
    }
  }),
  { name: "Skeleton" }
);

const Skeleton =
  decorate <
  StandardProps >
  (({ classes, className, style, ...props }) => (
    <span
      className={[classes.skeleton, className || ""].join(" ")}
      style={style}
    >
      &zwnj;
    </span>
  ));
export default Skeleton;
