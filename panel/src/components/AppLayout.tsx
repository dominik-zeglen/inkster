import * as React from "react";
import withStyles from "react-jss";
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

import IconButton from "./IconButton";
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

const decorate = withStyles(theme => ({
  active: {},
  content: {
    flexGrow: 1,
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
    marginBottom: theme.spacing,
    overflow: "hidden",
    padding: theme.spacing,
    transition: theme.transition.time,
  },
  menuText: {
    left: 0,
    position: "relative" as "relative",
  },
  root: {
    display: "flex" as "flex",
  },
  shrinkMenu: {
    display: "flex" as "flex",
    justifyContent: "center" as "center",
    marginBottom: theme.spacing,
  },
  sideMenu: {
    background: `linear-gradient(-45deg, ${theme.colors.black.main}, ${
      theme.colors.black.dark
    })`,
    color: theme.colors.gray.lightest,
    display: "flex" as "flex",
    flexDirection: "column" as "column",
    overflow: "hidden",
    padding: theme.spacing * 2,
    width: SIDEBAR_WIDTH,
    maxWidth: SIDEBAR_WIDTH,
    minHeight: "100vh",
    transition: theme.transition.time,
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
  spacer: {
    flex: 1,
  },
}));
export const AppLayout = decorate<Props>(
  ({ classes, children, section, onLogout, onSectionClick }) => (
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
            <IconButton
              className={classes.shrinkMenu}
              icon={isMenuShrunken ? Maximize2 : Minimize2}
              onClick={shrinkMenu}
            />
          </div>
          <div className={classes.content}>{children}</div>
        </div>
      )}
    </Toggle>
  ),
);
export default AppLayout;
