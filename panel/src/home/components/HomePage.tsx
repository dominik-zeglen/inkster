import * as React from "react";
import withStyles from "react-jss";

import Container from "../../components/Container";
import PageHeader from "../../components/PageHeader";
import i18n from "../../i18n";
import HomePageNewestPages from "./HomePageNewestPages";
import { Viewer_viewer } from "../queries/types/Viewer";
import { maybe } from "../../utils";

export interface Props {
  disabled: boolean;
  user: Viewer_viewer;
  onPageClick: (id: string) => void;
}

const decorate = withStyles(theme => ({
  root: {
    display: "grid",
    gridTemplateColumns: "1fr 1fr",
    gridColumnGap: theme.spacing * 2 + "px",
  },
}));

export const HomePage = decorate<Props>(
  ({ classes, disabled, user, onPageClick }) => (
    <Container width="md">
      <PageHeader
        title={
          user && user.email
            ? i18n.t("Hello, {{ email }}!", {
                context: "home page header",
                email: user.email,
              })
            : undefined
        }
      />
      <div className={classes.root}>
        <div>
          <HomePageNewestPages
            disabled={disabled}
            pages={maybe(() => user.pages.edges.map(edge => edge.node))}
            onPageClick={onPageClick}
          />
        </div>
      </div>
    </Container>
  ),
);
export default HomePage;
