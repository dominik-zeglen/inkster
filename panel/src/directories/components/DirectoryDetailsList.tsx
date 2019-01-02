import * as React from "react";
import { Panel } from "react-bootstrap";
import { Folder } from "react-feather";

import { ViewProps, PaginatedListProps } from "../..";
import ListElement from "../../components/ListElement";
import PaginationArrows from "../../components/PaginationArrows";
import i18n from "../../i18n";

interface Props extends ViewProps, PaginatedListProps {
  directories?: Array<{
    id: string;
    name?: string;
  }>;
}

export const DirectoryDetailsList: React.StatelessComponent<Props> = ({
  directories,
  disabled,
  loading,
  pageInfo,
  onNextPage,
  onPreviousPage,
  onRowClick,
}) => (
  <Panel>
    <Panel.Body>
      {directories ? (
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
export default DirectoryDetailsList;
