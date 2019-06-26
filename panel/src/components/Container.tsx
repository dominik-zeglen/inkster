import * as React from "react";
import createUseStyles from "aurora-ui-kit/dist/utils/jss";
import { ITheme } from "aurora-ui-kit/dist/theme";

type ContainerWidth = "xs" | "sm" | "md" | "lg" | "xl";
interface Props {
  width: ContainerWidth;
}

const useStyles = createUseStyles((theme: ITheme) =>
  (["xs", "sm", "md", "lg", "xl"] as ContainerWidth[]).reduce(
    (prev, current) => {
      prev[current] = {
        [theme.breakpoints.up(current)]: {
          marginLeft: "auto",
          marginRight: "auto",
          maxWidth: theme.breakpoints.width(current),
        },
      };
      return prev;
    },
    {},
  ),
);
export const Container: React.FC<Props> = ({ width, ...props }) => {
  const classes = useStyles();

  return <div className={classes[width]} {...props} />;
};
export default Container;
