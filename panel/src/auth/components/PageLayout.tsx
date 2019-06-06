import * as React from "react";
import withStyles from "react-jss";
import Card from "aurora-ui-kit/dist/components/Card";

import Container from "../../components/Container";

interface Props {
  header: string;
}

const decorate = withStyles({
  header: {
    textAlign: "center" as "center",
    textTransform: "uppercase" as "uppercase",
  },
  panel: {
    padding: 60,
    width: "100%",
  },
  root: {
    alignItems: "center" as "center",
    display: "flex" as "flex",
    height: "100vh",
    padding: 80,
  },
});
export const PageLayout = decorate<Props>(({ classes, children, header }) => (
  <Container width="xs">
    <div className={classes.root}>
      <Card className={classes.panel}>
        <h2 className={classes.header}>{header}</h2>
        {children}
      </Card>
    </div>
  </Container>
));
export default PageLayout;
