import * as React from "react";
import { Panel } from "react-bootstrap";
import { Folder } from "react-feather";

import { ViewProps, PaginatedListProps } from "../..";
import ListElement from "../../components/ListElement";
import PaginationArrows from "../../components/PaginationArrows";
import i18n from "../../i18n";
import { RootDirectories_getRootDirectories_edges_node } from "../queries/types/RootDirectories";

interface Props extends ViewProps, PaginatedListProps {
  directories: RootDirectories_getRootDirectories_edges_node[];
}

export const DirectoryRootList: React.StatelessComponent<Props> = ({
  directories,
  disabled,
  pageInfo,
  onNextPage,
  onPreviousPage,
  onRowClick,
}) => (
  <Panel>
    <Panel.Body>
      {directories !== undefined ? (
        directories.length > 0 ? (
          directories.map(directory => (
            <ListElement
              disabled={disabled}
              title={directory.name}
              onClick={onRowClick(directory.id)}
              icon={Folder}
            />
          ))
        ) : (
          i18n.t("No directories found")
        )
      ) : (
        <ListElement disabled={disabled} icon={Folder} />
      )}
    </Panel.Body>
    <Panel.Footer>
      <PaginationArrows
        pageInfo={pageInfo}
        onNextPage={onNextPage}
        onPreviousPage={onPreviousPage}
      />
    </Panel.Footer>
  </Panel>
);
export default DirectoryRootList;
