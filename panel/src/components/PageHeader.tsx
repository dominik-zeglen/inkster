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
    marginTop: theme.spacing,
    marginBottom: theme.spacing * 2
  },
  container: {
    alignItems: "center",
    display: "flex"
  },
  title: {
    flex: 1,
    marginLeft: theme.spacing,
    marginTop: theme.spacing / 2,
    ...theme.typography.mainHeading
  }
}));
export const PageHeader =
  decorate <
  Props >
  (({ classes, title, onBack }) => (
    <div className={classes.root}>
      <div className={classes.container}>
        {!!onBack ? <IconButton icon={ArrowLeft} onClick={onBack} /> : <div />}
        {title ? (
          <span className={classes.title}>{title}</span>
        ) : (
          <Skeleton
            className={classes.title}
            style={{ width: "14rem", flex: "unset" }}
          />
        )}
      </div>
    </div>
  ));
export default PageHeader;
