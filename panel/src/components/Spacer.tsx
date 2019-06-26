import * as React from "react";
import { ITheme } from "aurora-ui-kit/dist/theme";
import createUseStyles from "aurora-ui-kit/dist/utils/jss";

const useStyles = createUseStyles((theme: ITheme) => ({
  root: {
    marginBottom: theme.spacing * 2,
  },
}));
export const Spacer: React.FC = () => {
  const classes = useStyles();
  return <div className={classes.root} />;
};
Spacer.displayName = "Spacer";
export default Spacer;
