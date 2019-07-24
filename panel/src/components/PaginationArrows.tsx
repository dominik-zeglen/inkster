import * as React from "react";
import { ArrowLeft, ArrowRight } from "react-feather";
import IconButton from "aurora-ui-kit/dist/components/IconButton";
import createUseStyles from "aurora-ui-kit/dist/utils/jss";
import { ITheme } from "aurora-ui-kit/dist/theme";

interface Props {
  disabled: boolean;
  pageInfo?: {
    hasNextPage: boolean;
    hasPreviousPage: boolean;
  };
  onNextPage: () => void;
  onPreviousPage: () => void;
}

const useStyles = createUseStyles((theme: ITheme) => ({
  root: {
    display: "grid",
    gridColumnGap: theme.spacing + "px",
    gridTemplateColumns: `1fr 40px 40px`,
  },
}));

export const PaginationArrows: React.FC<Props> = ({
  disabled,
  pageInfo,
  onNextPage,
  onPreviousPage,
}) => {
  const classes = useStyles();
  return (
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
  );
};
export default PaginationArrows;
