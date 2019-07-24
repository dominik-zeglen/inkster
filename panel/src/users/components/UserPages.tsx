import * as React from "react";
import Card from "aurora-ui-kit/dist/components/Card";
import CardHeader from "aurora-ui-kit/dist/components/CardHeader";
import CardTitle from "aurora-ui-kit/dist/components/CardTitle";
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
import { maybe, renderCollection } from "../../utils";
import { UserDetails_user_pages_edges_node } from "../queries/types/UserDetails";
import Date from "../../components/Date";

interface Props extends ViewProps, PaginatedListProps {
  pages: UserDetails_user_pages_edges_node[];
}

const useStyles = createUseStyles({
  colStatus: css`
    text-align: center;
  `,
  row: {
    cursor: "pointer",
  },
});
export const UserPages: React.FC<Props> = ({
  disabled,
  pageInfo,
  pages,
  onNextPage,
  onPreviousPage,
  onRowClick,
}) => {
  const classes = useStyles();
  return (
    <Card>
      <CardHeader>
        <CardTitle>{i18n.t("Written by  user")}</CardTitle>
      </CardHeader>
      <Table>
        <TableHead>
          <TableCell>{i18n.t("Title")}</TableCell>
          <TableCell>{i18n.t("Created")}</TableCell>
          <TableCell className={classes.colStatus}>
            {i18n.t("Status")}
          </TableCell>
        </TableHead>
        <TableFooter>
          <TableRow>
            <TableCell colSpan={3}>
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
            pages,
            page => (
              <TableRow
                className={classes.row}
                hover={!!page}
                onClick={maybe(() => onRowClick(page.id))}
              >
                <TableCell>
                  {maybe<React.ReactNode>(() => page.name, <Skeleton />)}
                </TableCell>
                <TableCell>
                  {maybe<React.ReactNode>(
                    () => (
                      <Date date={page.createdAt} />
                    ),
                    <Skeleton />,
                  )}
                </TableCell>
                <TableCell className={classes.colStatus}>
                  {page && page.isPublished !== undefined ? (
                    <Status color={page.isPublished ? "primary" : "disabled"}>
                      {page.isPublished
                        ? i18n.t("Published")
                        : i18n.t("Unpublished")}
                    </Status>
                  ) : (
                    <Skeleton />
                  )}
                </TableCell>
              </TableRow>
            ),
            () => (
              <TableRow>
                <TableCell colSpan={3}>
                  {i18n.t("This user hasn't written anything yet")}
                </TableCell>
              </TableRow>
            ),
          )}
        </TableBody>
      </Table>
    </Card>
  );
};
export default UserPages;
