import * as React from "react";
import {
  Breadcrumb,
  Button,
  Col,
  Glyphicon,
  ListGroup,
  ListGroupItem,
  Panel,
  Row
} from "react-bootstrap";
import withStyles from "react-jss";

interface ContainerListProps {
  containers?: Array<{
    id: number;
    name: string;
  }>;
  loading?: boolean;
  onRowClick(id);
}

const decorate = withStyles(theme => ({
  button: theme.components.button,
  flexStretch: {
    flex: 1
  },
  iconButton: {
    ...theme.components.iconButton
  },
  listGroupItem: theme.components.listGroupItem,
  panel: theme.components.panel
}));

export const ContainerList = decorate<ContainerListProps>(
  ({ containers, classes, onRowClick, loading }) => (
    <Row>
      <Col xs={12} md={8}>
        <Panel className={classes.panel}>
          <Panel.Heading>
            <span>Containers</span>
            <div className={classes.flexStretch} />
            <Button className={classes.iconButton}>
              <Glyphicon glyph="plus" />
            </Button>
          </Panel.Heading>
          <ListGroup>
            {loading ? (
              <span>loading</span>
            ) : containers.length > 0 ? (
              containers.map(container => (
                <ListGroupItem
                  onClick={onRowClick(container.id)}
                  className={classes.listGroupItem}
                  key={container.id}
                >
                  {container.name}
                </ListGroupItem>
              ))
            ) : (
              <ListGroupItem>Sry matey no containers here</ListGroupItem>
            )}
          </ListGroup>
        </Panel>
      </Col>
    </Row>
  )
);

export default ContainerList;
