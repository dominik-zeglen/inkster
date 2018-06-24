import * as React from "react";
import { Panel } from "react-bootstrap";
import withStyles from "react-jss";
import { Box } from "react-feather";

import Paginator from "../../components/Paginator";
import Skeleton from "../../components/Skeleton";
import i18n from "../../i18n";

interface ListProps {
  containers?: Array<{
    id?: string,
    name?: string
  }>;
  hasNextPage?: boolean;
  hasPreviousPage?: boolean;
  onNextPage?: () => void;
  onPreviousPage?: () => void;
  onRowClick?: (id: string) => () => void;
}
interface TileProps {
  loading?: boolean;
  name?: string;
  onClick?: () => void;
}

const decorate = withStyles((theme: any) => ({
  root: {
    cursor: "pointer",
    display: "grid",
    gridColumnGap: theme.spacing + "px",
    gridTemplateColumns: `${theme.spacing * 5}px 1fr`,
    marginTop: theme.spacing,
    marginBottom: theme.spacing,
    transitionDuration: theme.transition.time,
    "&:hover, &:focus": {
      color: theme.colors.secondary.dark,
      "& svg": {
        color: theme.colors.secondary.main
      }
    }
  }
}));
const ContainerTile =
  decorate <
  TileProps >
  (({ classes, loading, name, onClick }) => (
    <div className={classes.root} onClick={loading ? undefined : onClick}>
      <Box />
      <div>
        {loading || !name ? <Skeleton style={{ width: "10rem" }} /> : name}
      </div>
    </div>
  ));

export const ContainerListPage: React.StatelessComponent<ListProps> = ({
  containers,
  hasNextPage,
  hasPreviousPage,
  onNextPage,
  onPreviousPage,
  onRowClick
}) => (
  <Panel>
    <Panel.Body>
      {containers === undefined ? (
        <ContainerTile loading={true} />
      ) : containers.length > 0 ? (
        containers.map(container => (
          <ContainerTile
            onClick={!!onRowClick && container.id ? onRowClick(container.id) : undefined}
            name={container.name}
          />
        ))
      ) : (
        i18n.t(
          'No containers found. You can add one by clicking "Plus" button on the top right corner of the page.'
        )
      )}
    </Panel.Body>
    <Panel.Footer>
      <Paginator
        hasNextPage={hasNextPage}
        hasPreviousPage={hasPreviousPage}
        onPreviousPage={onPreviousPage}
        onNextPage={onNextPage}
      />
    </Panel.Footer>
  </Panel>
);
export default ContainerListPage;
