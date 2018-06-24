import * as React from "react";
import { ArrowLeft, ArrowRight } from "react-feather";
import withStyles from "react-jss";

import IconButton from "./IconButton";

interface Props {
  hasNextPage?: boolean;
  hasPreviousPage?: boolean;
  onNextPage?: () => void;
  onPreviousPage?: () => void;
}

const decorate = withStyles((theme: any) => ({
  root: {
    display: "grid",
    gridColumnGap: theme.spacing + "px",
    gridTemplateColumns: `1fr ${theme.spacing * 2.5}px ${theme.spacing * 2.5}px`
  }
}));

export const Paginator =
  decorate <
  Props >
  (({ classes, hasNextPage, hasPreviousPage, onNextPage, onPreviousPage }) => (
    <div className={classes.root}>
      <div />
      <IconButton
        icon={ArrowLeft}
        disabled={!(!!onPreviousPage && hasPreviousPage)}
        onClick={onPreviousPage}
      />
      <IconButton
        icon={ArrowRight}
        disabled={!(!!onNextPage && hasNextPage)}
        onClick={onNextPage}
      />
    </div>
  ));
export default Paginator;
