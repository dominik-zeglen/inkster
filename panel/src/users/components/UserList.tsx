import * as React from "react";
import { Panel } from "react-bootstrap";
import { User as UserIcon } from "react-feather";

import i18n from '../../i18n'
import ListElement from "../../components/ListElement";
import Paginator from "../../components/Paginator";
import { ViewProps, PaginatedListProps } from "../..";

interface Props extends ViewProps, PaginatedListProps {
  users?: Array<{
    id: string;
    email: string;
  }>
}

export const UserListPage: React.StatelessComponent<Props> = ({
  disabled,
  loading,
  pageInfo,
  users,
  onNextPage,
  onPreviousPage,
  onRowClick
}) => (
  <Panel>
    <Panel.Body>
      {users !== undefined ? (
        users.length > 0 ? (
          users.map(user => (
            <ListElement
              disabled={disabled}
              title={user.email}
              onClick={onRowClick(user.id)}
              icon={UserIcon}
            />
          ))
        ) : (
          i18n.t("No users found")
        )
      ) : (
        <ListElement disabled={disabled} icon={UserIcon} />
      )}
    </Panel.Body>
    <Panel.Footer>
      <Paginator
        pageInfo={pageInfo}
        onNextPage={onNextPage}
        onPreviousPage={onPreviousPage}
      />
    </Panel.Footer>
  </Panel>
);
export default UserListPage;
