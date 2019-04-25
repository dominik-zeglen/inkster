import * as React from "react";
import withStyles from "react-jss";

const decorate = withStyles(theme => ({
  root: {
    marginBottom: theme.spacing * 2,
  },
}));
export const Spacer = decorate(({ classes }) => (
  <div className={classes.root} />
));
Spacer.displayName = "Spacer";
export default Spacer;
