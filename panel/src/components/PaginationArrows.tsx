import * as React from "react";
import { ArrowLeft, ArrowRight } from "react-feather";
import withStyles from "react-jss";
import IconButton from "aurora-ui-kit/dist/components/IconButton";

interface Props {
  disabled: boolean;
  pageInfo?: {
    hasNextPage: boolean;
    hasPreviousPage: boolean;
  };
  onNextPage: () => void;
  onPreviousPage: () => void;
}

const decorate = withStyles((theme: any) => ({
  root: {
    display: "grid",
    gridColumnGap: theme.spacing * 3 + "px",
    gridTemplateColumns: `1fr ${theme.spacing * 2.5}px ${theme.spacing *
      2.5}px`,
    paddingRight: theme.spacing * 2,
  },
}));

export const PaginationArrows = decorate<Props>(
  ({ classes, disabled, pageInfo, onNextPage, onPreviousPage }) => (
    <div className={classes.root}>
      <div />
      <IconButton
        disabled={
          !(!!onPreviousPage && pageInfo && pageInfo.hasPreviousPage) ||
          disabled
        }
        onClick={onPreviousPage}
      >
        <ArrowLeft />
      </IconButton>
      <IconButton
        disabled={
          !(!!onNextPage && pageInfo && pageInfo.hasNextPage) || disabled
        }
        onClick={onNextPage}
      >
        <ArrowRight />
      </IconButton>
    </div>
  ),
);
export default PaginationArrows;
