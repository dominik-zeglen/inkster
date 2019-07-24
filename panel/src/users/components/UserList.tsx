import * as React from "react";
import Card from "aurora-ui-kit/dist/components/Card";
import Table from "aurora-ui-kit/dist/components/Table";
import TableRow from "aurora-ui-kit/dist/components/TableRow";
import TableCell from "aurora-ui-kit/dist/components/TableCell";
import TableBody from "aurora-ui-kit/dist/components/TableBody";
import TableHead from "aurora-ui-kit/dist/components/TableHead";
import TableFooter from "aurora-ui-kit/dist/components/TableFooter";
import Skeleton from "aurora-ui-kit/dist/components/Skeleton";
import Status from "aurora-ui-kit/dist/components/Status";
import createUseStyles, { css } from "aurora-ui-kit/dist/utils/jss";

import i18n from "../../i18n";
import PaginationArrows from "../../components/PaginationArrows";
import { ViewProps, PaginatedListProps } from "../..";
import { UserList_users_edges_node } from "../queries/types/UserList";
import { maybe, renderCollection } from "../../utils";

interface Props extends ViewProps, PaginatedListProps {
  users: UserList_users_edges_node[];
}

const useStyles = createUseStyles({
  colStatus: css`
    text-align: center;
  `,
  row: {
    cursor: "pointer",
  },
});
export const UserListPage: React.FC<Props> = ({
  disabled,
  pageInfo,
  users,
  onNextPage,
  onPreviousPage,
  onRowClick,
}) => {
  const classes = useStyles();
  return (
    <Card>
      <Table>
        <TableHead>
          <TableCell>{i18n.t("E-mail Address")}</TableCell>
          <TableCell className={classes.colStatus}>
            {i18n.t("Status")}
          </TableCell>
        </TableHead>
        <TableFooter>
          <TableRow>
            <TableCell colSpan={2}>
              <PaginationArrows
                disabled={disabled}
                pageInfo={pageInfo}
                onNextPage={onNextPage}
                onPreviousPage={onPreviousPage}
              />
            </TableCell>
          </TableRow>
        </TableFooter>
        <TableBody>
          {renderCollection(
            users,
            user => (
              <TableRow
                className={classes.row}
                hover={!!user}
                onClick={maybe(() => onRowClick(user.id))}
              >
                <TableCell>
                  {maybe<React.ReactNode>(() => user.email, <Skeleton />)}
                </TableCell>
                <TableCell className={classes.colStatus}>
                  {user && user.isActive !== undefined ? (
                    <Status color={user.isActive ? "primary" : "disabled"}>
                      {user.isActive ? i18n.t("Active") : i18n.t("Inactive")}
                    </Status>
                  ) : (
                    <Skeleton />
                  )}
                </TableCell>
              </TableRow>
            ),
            () => (
              <TableRow>
                <TableCell colSpan={2}>{i18n.t("No users found")}</TableCell>
              </TableRow>
            ),
          )}
        </TableBody>
      </Table>
    </Card>
  );
};
export default UserListPage;
