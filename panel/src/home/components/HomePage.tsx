import * as React from "react";

import Container from "../../components/Container";
import PageHeader from "../../components/PageHeader";
import i18n from "../../i18n";
import HomePageNewestPages from "./HomePageNewestPages";
import { Viewer_viewer } from "../queries/types/Viewer";
import { maybe } from "../../utils";
import { ITheme } from "aurora-ui-kit/dist/theme";
import createUseStyles from "aurora-ui-kit/dist/utils/jss";

export interface Props {
  disabled: boolean;
  user: Viewer_viewer;
  onPageClick: (id: string) => void;
}

const useStyles = createUseStyles((theme: ITheme) => ({
  root: {
    display: "grid",
    gridTemplateColumns: "1fr 1fr",
    gridColumnGap: theme.spacing * 2 + "px",
  },
}));

export const HomePage: React.FC<Props> = ({ disabled, user, onPageClick }) => {
  const classes = useStyles();
  return (
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
  );
};
export default HomePage;
