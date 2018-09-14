import * as React from "react";
import withStyles from "react-jss";
import { Grid, Panel } from "react-bootstrap";
import { Box, Home, Users } from "react-feather";

import i18n from "../i18n";

interface Props {
  section: string;
  onSectionClick: (section: string) => () => void;
}

const decorate = withStyles((theme: any) => ({
  link: {
    '&.active': {
      color: theme.colors.secondary.dark
    },
    "&:hover": {
      color: theme.colors.secondary.main
    },
    "& svg": {
      marginRight: theme.spacing
    },
    alignItems: "center" as "center",
    cursor: "pointer" as "pointer",
    display: "flex" as "flex",
    marginBottom: theme.spacing,
    transition: theme.transition.time
  },
  navbar: {
    background: theme.colors.primary.light,
    boxShadow: "0px 5px 20px 5px #f2f2f2",
    height: theme.spacing * 6,
    marginBottom: theme.spacing * 2,
    width: "100%"
  },
  root: {
    display: "grid" as "grid",
    gridColumnGap: theme.spacing + "px",
    gridTemplateColumns: `${theme.spacing * 25}px 1fr`
  },
  sideMenu: {
    marginTop: theme.spacing * 7.3,
    width: 22 * theme.spacing
  }
}));
export const AppLayout = decorate<Props>(
  ({ classes, children, section, onSectionClick }) => (
    <>
      <Grid>
        <div className={classes.root}>
          <div>
            <Panel className={classes.sideMenu}>
              <Panel.Heading>
                <Panel.Title>{i18n.t("Navigation")}</Panel.Title>
              </Panel.Heading>
              <Panel.Body>
                <div
                  className={[
                    classes.link,
                    section === "home" ? "active" : undefined
                  ].join(' ')}
                  onClick={onSectionClick("home")}
                >
                  <Home />
                  {i18n.t("Home")}
                </div>
                <div
                  className={[
                    classes.link,
                    section === "directories" ? "active" : undefined
                  ].join(' ')}
                  onClick={onSectionClick("directories")}
                >
                  <Box />
                  {i18n.t("Directories")}
                </div>
                <div
                  className={[
                    classes.link,
                    section === "users" ? "active" : undefined
                  ].join(' ')}
                  onClick={onSectionClick("users")}
                >
                  <Users />
                  {i18n.t("Users")}
                </div>
              </Panel.Body>
            </Panel>
          </div>
          <div>{children}</div>
        </div>
      </Grid>
    </>
  )
);
export default AppLayout;
