import * as React from "react";
import withStyles from "react-jss";

import { StandardProps } from ".";
import Skeleton from "./Skeleton";

interface Props extends StandardProps {
  disabled: boolean;
  icon: React.StatelessComponent<{ color?: string } & StandardProps>;
  title?: string;
}

const decorate = withStyles(
  (theme: any) => ({
    root: {
      display: "grid" as "grid",
      gridColumnGap: theme.spacing + "px",
      gridTemplateColumns: "2.5rem 1fr",
      marginBottom: theme.spacing / 2,
      marginLeft: theme.spacing / 2,
      marginTop: theme.spacing * 1.5,
      transition: theme.transition.time,
      "&:first-of-type": {
        marginTop: 0
      }
    },
    title: {
      marginTop: 0
    },
    link: {
      cursor: "pointer" as "pointer",
      "&:hover, &:focus": {
        color: theme.colors.secondary.main
      }
    }
  }),
  { displayName: "ListElement" }
);
export const ListElement = decorate<Props>(
  ({ className, classes, disabled, icon, title, onClick, ...props }) => {
    const Icon = icon;
    return (
      <div
        className={
          disabled
            ? [classes.root, className].join(" ")
            : [classes.root, classes.link, className].join(" ")
        }
        onClick={onClick}
        {...props}
      >
        <Icon />
        <div className={classes.title}>{title ? title : <Skeleton />}</div>
      </div>
    );
  }
);
export default ListElement;
