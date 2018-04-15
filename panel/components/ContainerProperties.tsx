import * as React from "react";
import {
  Breadcrumb,
  Button,
  Col,
  FormControl,
  Glyphicon,
  ListGroup,
  ListGroupItem,
  Panel,
  Row
} from "react-bootstrap";
import withStyles from "react-jss";

interface ContainerListProps {
  container?: {
    id: number;
    name: string;
  };
  loading?: boolean;
  onEdit(event: any);
  onRemove();
}

const decorate = withStyles(theme => ({
  button: {
    ...theme.components.button,
    backgroundColor: theme.colors.default
  },
  flexStretch: {
    flex: 1
  },
  iconButton: {
    ...theme.components.iconButton
  },
  panel: theme.components.panel
}));

export const ContainerList = decorate<ContainerListProps>(
  ({ container, classes, loading, onSubmit }) => (
    <Row>
      <Col xs={12} md={8}>
        <form onSubmit={onSubmit}>
          <Panel className={classes.panel}>
            <Panel.Heading>
              <span>{loading ? "" : container.name}</span>
              <div className={classes.flexStretch} />
              <Button className={classes.iconButton}>
                <Glyphicon glyph="trash" />
              </Button>
            </Panel.Heading>
            <Panel.Body>
              <FormControl
                name="name"
                defaultValue={loading ? "" : container.name}
                placeholder="Container name"
                type="text"
              />
            </Panel.Body>
            <Panel.Footer>
              <div className={classes.flexStretch} />
              <Button
                bsStyle="primary"
                onClick={onSubmit}
                className={classes.button}
              >
                Submit
              </Button>
            </Panel.Footer>
          </Panel>
        </form>
      </Col>
    </Row>
  )
);

export default ContainerList;
