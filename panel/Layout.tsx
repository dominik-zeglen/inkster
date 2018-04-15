import * as React from "react";
import {
  Col,
  FormControl,
  FormGroup,
  Grid,
  ListGroup,
  ListGroupItem,
  Navbar,
  NavbarBrand,
  Panel,
  Row
} from "react-bootstrap";
import withStyles from "react-jss";

interface LayoutProps {
  onHomeClick();
  onContainersClick();
  onBrandClick?();
  onSearchSubmit?(event: any);
}

const decorate = withStyles(theme => ({
  brand: {
    "&:not(\\20):not(#\\20)": {
      color: "#ffffff",
      cursor: "pointer"
    }
  },
  form: {
    margin: 0
  },
  listGroupItem: theme.components.listGroupItem,
  navbar: {
    "&:not(\\20)": {
      background: theme.colors.gradient
    }
  },
  panel: {
    ...theme.components.panel
  },
  search: {
    "&:not(\\20):not(#\\20)": {
      ...theme.borders,
      "&::placeholder": {
        color: "#ffffff"
      },
      "&:focus": {
        "&::placeholder": {
          color: theme.typography.disabled.color
        },
        background: "#fff",
        color: theme.typography.default.color
      },
      background: "transparent",
      boxShadow: "none",
      color: "#ffffff",
      marginTop: ".75rem",
      transition: "200ms"
    }
  }
}));

export const Layout = decorate<LayoutProps>(
  ({ children, classes, onBrandClick, onSearchSubmit, ...props }) => (
    <>
      <Navbar staticTop className={classes.navbar}>
        <Navbar.Header>
          <Row>
            <Grid>
              <Row>
                <Col xs={12} md={3} lg={2}>
                  <NavbarBrand className={classes.brand} onClick={onBrandClick}>
                    FOXXY
                  </NavbarBrand>
                </Col>
                <Col xs={12} md={7}>
                  <form className={classes.form} onSubmit={onSearchSubmit}>
                    <FormControl
                      name="search"
                      type="text"
                      placeholder="Type to search..."
                      className={classes.search}
                    />
                  </form>
                </Col>
              </Row>
            </Grid>
          </Row>
        </Navbar.Header>
      </Navbar>
      <Grid>
        <Row>
          <Col md={3} lg={2}>
            <Panel className={classes.panel}>
              <Panel.Heading>Menu</Panel.Heading>
              <ListGroup>
                <ListGroupItem
                  onClick={props.onHomeClick}
                  className={classes.listGroupItem}
                >
                  Home
                </ListGroupItem>
                <ListGroupItem
                  onClick={props.onContainersClick}
                  className={classes.listGroupItem}
                >
                  Containers
                </ListGroupItem>
              </ListGroup>
            </Panel>
          </Col>
          <Col md={9} lg={10}>
            {children}
          </Col>
        </Row>
      </Grid>
    </>
  )
);
export default Layout;
