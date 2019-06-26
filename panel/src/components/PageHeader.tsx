import * as React from "react";
import { ArrowLeft } from "react-feather";
import IconButton from "aurora-ui-kit/dist/components/IconButton";

import { ITheme } from "aurora-ui-kit/dist/theme";
import createUseStyles from "aurora-ui-kit/dist/utils/jss";
import { getFontSize } from "aurora-ui-kit/dist/components/Typography";
import Skeleton from "aurora-ui-kit/dist/components/Skeleton";

interface Props {
  title?: string | React.ReactNode;
  onBack?: () => void;
}

const useStyles = createUseStyles((theme: ITheme) => ({
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
    fontSize: getFontSize("mainHeading"),
  },
}));
export const PageHeader: React.FC<Props> = ({ children, title, onBack }) => {
  const classes = useStyles();
  return (
    <div className={classes.root}>
      <div className={classes.container}>
        {!!onBack ? (
          <IconButton onClick={onBack}>
            <ArrowLeft />
          </IconButton>
        ) : null}
        <span className={classes.title}>
          {title ? title : <Skeleton style={{ width: "14rem" }} />}
        </span>
        {children}
      </div>
    </div>
  );
};
export default PageHeader;
