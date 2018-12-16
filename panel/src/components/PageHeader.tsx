import * as React from "react";
import withStyles from "react-jss";
import { ArrowLeft } from "react-feather";

import IconButton from "./IconButton";
import Skeleton from "./Skeleton";

interface Props {
  title?: string | React.ReactNode;
  onBack?: () => void;
}

const decorate = withStyles((theme: any) => ({
  root: {
    marginTop: theme.spacing * 2.5,
    marginBottom: theme.spacing * 2,
  },
  container: {
    alignItems: "center",
    display: "flex",
  },
  title: {
    flex: 1,
    marginLeft: theme.spacing,
    marginTop: -2,
    ...theme.typography.mainHeading,
  },
}));
export const PageHeader = decorate<Props>(
  ({ children, classes, title, onBack }) => (
    <div className={classes.root}>
      <div className={classes.container}>
        {!!onBack ? <IconButton icon={ArrowLeft} onClick={onBack} /> : <div />}
        <span className={classes.title}>
          {title ? title : <Skeleton style={{ width: "14rem" }} />}
        </span>
        {children}
      </div>
    </div>
  ),
);
export default PageHeader;
