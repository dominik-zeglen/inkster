import * as React from "react";
import withStyles from "react-jss";
import Progress from "aurora-ui-kit/dist/components/LinearProgress";
import Typography from "aurora-ui-kit/dist/components/Typography";

interface Props {
  progress: number;
}

const decorate = withStyles((theme: any) => ({
  center: {
    display: "flex",
    flexDirection: "column" as "column",
    alignItems: "center",
    justifyContent: "center",
    height: "100%",
  },
  overlay: {
    position: "fixed" as "fixed",
    top: 0,
    left: 0,
    backgroundColor: "rgba(0, 0, 0, 0.1)",
    zIndex: 100,
    width: "100vw",
    height: "100vh",
  },
  root: {
    display: "flex" as "flex",
    alignItems: "center" as "center",
  },
  progressBar: {
    width: theme.spacing * 30,
  },
}));

export const LoaderOverlay = decorate<Props>(({ classes, progress }) => (
  <div className={classes.overlay}>
    <Progress />
    <div className={classes.center}>
      <div className={classes.root}>
        <Typography variant="mainHeading">
          {progress.toLocaleString("en", {
            style: "percent",
          })}
        </Typography>
      </div>
    </div>
  </div>
));
export default LoaderOverlay;
