import * as React from "react";
import createUseStyles, { css } from "aurora-ui-kit/dist/utils/jss";
import {
  Box,
  Home,
  LogOut,
  Maximize2,
  Minimize2,
  Users,
  Settings,
} from "react-feather";
import * as classNames from "classnames";
import IconButton from "aurora-ui-kit/dist/components/IconButton";
import { ITheme } from "aurora-ui-kit/dist/theme";

import Toggle from "./Toggle";
import i18n from "../i18n";

const SIDEBAR_WIDTH = 210;
const SIDEBAR_WIDTH_SHRUNKEN = 68;

interface Props {
  section: string;
  user: {
    id: string;
    email: string;
  };
  onLogout: () => void;
  onSectionClick: (section: string) => () => void;
}

const useStyles = createUseStyles((theme: ITheme) => ({
  active: {},
  content: {
    flexGrow: 1,
    padding: `0 ${theme.spacing * 3}px`,
  },
  link: {
    "&$active": {
      borderColor: theme.colors.secondary.dark,
      color: theme.colors.secondary.dark,
    },
    "&:hover": {
      color: theme.colors.secondary.main,
    },
    "& svg": {
      marginRight: theme.spacing * 2,
    },
    alignItems: "center" as "center",
    border: "2px solid transparent",
    borderRadius: 5,
    cursor: "pointer" as "pointer",
    display: "flex" as "flex",
    height: 48,
    marginBottom: theme.spacing,
    overflow: "hidden",
    padding: theme.spacing,
    transition: theme.transition.default,
  },
  menuText: {
    left: 0,
    position: "relative" as "relative",
  },
  root: {
    display: "flex" as "flex",
  },
  shrinkMenu: {
    alignItems: "center",
    color: theme.colors.gray.lightest,
    display: "flex" as "flex",
    height: 48,
    justifyContent: "center" as "center",
    marginBottom: theme.spacing,
    width: 48,
  },
  sideMenu: {
    background: `linear-gradient(45deg, ${theme.colors.gray.darkest}, ${
      theme.colors.common.black
    })`,
    color: theme.colors.gray.lightest,
    display: "flex" as "flex",
    flexDirection: "column" as "column",
    overflow: "hidden",
    padding: theme.spacing * 2,
    width: SIDEBAR_WIDTH,
    maxWidth: SIDEBAR_WIDTH,
    minHeight: "100vh",
    transition: theme.transition.default,
  },
  sideMenuShrunken: {
    "& $link": {
      "& $menuText": {
        maxWidth: 0,
        left: 300,
      },
      "& svg": {
        marginRight: 0,
      },
      borderColor: "transparent",
    },
    padding: `${theme.spacing * 2}px ${theme.spacing}px`,
    width: SIDEBAR_WIDTH_SHRUNKEN,
  },
  shrinkMenuIconContainer: css`
    display: flex;
    justify-content: center;
  `,
  spacer: {
    flex: 1,
  },
}));
export const AppLayout: React.FC<Props> = ({
  children,
  section,
  onLogout,
  onSectionClick,
}) => {
  const classes = useStyles();

  return (
    <Toggle initial={true}>
      {(isMenuShrunken, { toggle: shrinkMenu }) => (
        <div className={classes.root}>
          <div
            className={classNames({
              [classes.sideMenu]: true,
              [classes.sideMenuShrunken]: isMenuShrunken,
            })}
          >
            <div
              className={classNames({
                [classes.link]: true,
                [classes.active]: section === "home",
              })}
              onClick={onSectionClick("home")}
            >
              <Home />
              <span className={classes.menuText}>{i18n.t("Home")}</span>
            </div>
            <div
              className={classNames({
                [classes.link]: true,
                [classes.active]:
                  section === "directories" || section === "pages",
              })}
              onClick={onSectionClick("directories")}
            >
              <Box />
              <span className={classes.menuText}>{i18n.t("Directories")}</span>
            </div>
            <div
              className={classNames({
                [classes.link]: true,
                [classes.active]: section === "users",
              })}
              onClick={onSectionClick("users")}
            >
              <Users />
              <span className={classes.menuText}>{i18n.t("Users")}</span>
            </div>
            <div className={classes.spacer} />
            <div
              className={classNames({
                [classes.link]: true,
                [classes.active]: section === "settings",
              })}
              onClick={onSectionClick("settings")}
            >
              <Settings />
              <span className={classes.menuText}>{i18n.t("Settings")}</span>
            </div>
            <div className={classes.link} onClick={onLogout}>
              <LogOut />
              <span className={classes.menuText}>
                {i18n.t("Log out", {
                  context: "button",
                })}
              </span>
            </div>
            <div className={classes.shrinkMenuIconContainer}>
              <IconButton className={classes.shrinkMenu} onClick={shrinkMenu}>
                {isMenuShrunken ? <Maximize2 /> : <Minimize2 />}
              </IconButton>
            </div>
          </div>
          <div className={classes.content}>{children}</div>
        </div>
      )}
    </Toggle>
  );
};
export default AppLayout;
