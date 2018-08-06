import * as React from "react";
import withStyles from "react-jss";
import { ProgressBar } from "react-bootstrap";

interface Props {
  progress: number;
}

const decorate = withStyles((theme: any) => ({
  overlay: {
    position: "fixed" as "fixed",
    top: 0,
    left: 0,
    backgroundColor: "rgba(0, 0, 0, 0.1)",
    zIndex: 100,
    width: "100vw",
    height: "100vh"
  },
  root: {
    display: "flex" as "flex",
    alignItems: "center" as "center"
  },
  progressBar: {
    width: theme.spacing * 30
  }
}));

export const LoaderOverlay = decorate<Props>(({ classes, progress }) => (
  <div className={classes.overlay}>
    <div className={classes.root}>
      <ProgressBar now={progress} />
    </div>
  </div>
));
export default LoaderOverlay;
