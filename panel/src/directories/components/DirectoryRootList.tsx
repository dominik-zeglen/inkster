import * as React from "react";
import TableRow from "aurora-ui-kit/dist/components/TableRow";
import TableBody from "aurora-ui-kit/dist/components/TableBody";
import Skeleton from "aurora-ui-kit/dist/components/Skeleton";
import Card from "aurora-ui-kit/dist/components/Card";
import TableFooter from "aurora-ui-kit/dist/components/TableFooter";
import CardHeader from "aurora-ui-kit/dist/components/CardHeader";
import Status from "aurora-ui-kit/dist/components/Status";
import CardTitle from "aurora-ui-kit/dist/components/CardTitle";
import Table from "aurora-ui-kit/dist/components/Table";
import TableHead from "aurora-ui-kit/dist/components/TableHead";
import TableCell from "aurora-ui-kit/dist/components/TableCell";

import { ViewProps, PaginatedListProps } from "../..";
import PaginationArrows from "../../components/PaginationArrows";
import i18n from "../../i18n";
import { RootDirectories_getRootDirectories_edges_node } from "../queries/types/RootDirectories";
import { maybe, renderCollection } from "../../utils";
import createUseStyles, { css } from "aurora-ui-kit/dist/utils/jss";

interface Props extends ViewProps, PaginatedListProps {
  directories: RootDirectories_getRootDirectories_edges_node[];
}

const useStyles = createUseStyles({
  colStatus: css`
    text-align: center;
  `,
  row: {
    cursor: "pointer",
  },
});
export const DirectoryRootList: React.FC<Props> = ({
  directories,
  disabled,
  pageInfo,
  onNextPage,
  onPreviousPage,
  onRowClick,
}) => {
  const classes = useStyles();
  return (
    <Card>
      <CardHeader>
        <CardTitle>{i18n.t("Directories")}</CardTitle>
      </CardHeader>
      <Table>
        <TableHead>
          <TableCell>{i18n.t("Name")}</TableCell>
          <TableCell className={classes.colStatus}>
            {i18n.t("Status")}
          </TableCell>
        </TableHead>
        <TableFooter>
          <TableCell colSpan={100}>
            <PaginationArrows
              disabled={disabled}
              pageInfo={pageInfo}
              onNextPage={onNextPage}
              onPreviousPage={onPreviousPage}
            />
          </TableCell>
        </TableFooter>
        <TableBody>
          {renderCollection(
            directories,
            directory => (
              <TableRow
                className={classes.row}
                hover={!disabled}
                onClick={directory ? onRowClick(directory.id) : undefined}
              >
                <TableCell>
                  {maybe<React.ReactNode>(() => directory.name, <Skeleton />)}
                </TableCell>
                <TableCell className={classes.colStatus}>
                  {maybe(
                    () =>
                      directory.isPublished ? (
                        <Status color="primary">{i18n.t("Published")}</Status>
                      ) : (
                        <Status color="disabled">{i18n.t("Unublished")}</Status>
                      ),
                    <Skeleton />,
                  )}
                </TableCell>
              </TableRow>
            ),
            () => (
              <TableRow>
                <TableCell colSpan={100}>
                  {i18n.t("No directories found")}
                </TableCell>
              </TableRow>
            ),
          )}
        </TableBody>
      </Table>
    </Card>
  );
};
export default DirectoryRootList;
