import * as React from "react";
import { Panel } from "react-bootstrap";
import withStyles from "react-jss";
import { FileText } from "react-feather";

import Paginator from "../../components/Paginator";
import Skeleton from "../../components/Skeleton";
import i18n from "../../i18n";

interface ListProps {
  pages?: Array<{
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
    marginLeft: theme.spacing,
    transitionDuration: theme.transition.time,
    "&:hover, &:focus": {
      color: theme.colors.secondary.dark,
      "& svg": {
        color: theme.colors.secondary.main
      }
    }
  }
}));
const PageTile =
  decorate <
  TileProps >
  (({ classes, loading, name, onClick }) => (
    <div className={classes.root} onClick={loading ? undefined : onClick}>
      <FileText />
      <div>
        {loading || !name ? <Skeleton style={{ width: "10rem" }} /> : name}
      </div>
    </div>
  ));

export const PageList: React.StatelessComponent<ListProps> = ({
  pages,
  hasNextPage,
  hasPreviousPage,
  onNextPage,
  onPreviousPage,
  onRowClick
}) => (
  <Panel>
    <Panel.Body>
      {pages === undefined ? (
        <PageTile loading={true} />
      ) : pages.length > 0 ? (
        pages.map(page => (
          <PageTile
            onClick={!!onRowClick && page.id ? onRowClick(page.id) : undefined}
            name={page.name}
          />
        ))
      ) : (
        i18n.t(
          'No pages found. You can add one by clicking "Plus" button on the top right corner of the page.'
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
