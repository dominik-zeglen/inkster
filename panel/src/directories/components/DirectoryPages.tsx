import * as React from "react";
import { Plus } from "react-feather";
import TableRow from "aurora-ui-kit/dist/components/TableRow";
import withStyles, { CSSProperties } from "react-jss";
import TableBody from "aurora-ui-kit/dist/components/TableBody";
import Skeleton from "aurora-ui-kit/dist/components/Skeleton";
import IconButton from "aurora-ui-kit/dist/components/IconButton";
import Card from "aurora-ui-kit/dist/components/Card";
import TableFooter from "aurora-ui-kit/dist/components/TableFooter";
import CardHeader from "aurora-ui-kit/dist/components/CardHeader";
import CardTitle from "aurora-ui-kit/dist/components/CardTitle";
import Table from "aurora-ui-kit/dist/components/Table";
import TableHead from "aurora-ui-kit/dist/components/TableHead";
import TableCell from "aurora-ui-kit/dist/components/TableCell";
import Status from "aurora-ui-kit/dist/components/Status";

import { PaginatedListProps } from "../..";
import PaginationArrows from "../../components/PaginationArrows";
import i18n from "../../i18n";
import { Directory_getDirectory_pages_edges_node } from "../queries/types/Directory";
import { renderCollection, maybe } from "../../utils";

interface Props extends PaginatedListProps {
  disabled: boolean;
  pages: Directory_getDirectory_pages_edges_node[];
  onAdd: () => void;
}

const styles: CSSProperties = {
  row: {
    cursor: "pointer",
  },
};
export const DirectoryRootList = withStyles(styles)<Props>(
  ({
    classes,
    disabled,
    pages,
    pageInfo,
    onAdd,
    onNextPage,
    onPreviousPage,
    onRowClick,
  }) => (
    <Card>
      <CardHeader>
        <CardTitle>{i18n.t("Pages")}</CardTitle>
        <IconButton disabled={disabled} onClick={onAdd}>
          <Plus />
        </IconButton>
      </CardHeader>
      <Table>
        <TableHead>
          <TableCell>{i18n.t("Title")}</TableCell>
          <TableCell>{i18n.t("Status")}</TableCell>
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
            pages,
            page => (
              <TableRow
                className={classes.row}
                hover={!disabled}
                onClick={page ? onRowClick(page.id) : undefined}
              >
                <TableCell>
                  {maybe<React.ReactNode>(() => page.name, <Skeleton />)}
                </TableCell>
                <TableCell>
                  {maybe(
                    () =>
                      page.isPublished ? (
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
                <TableCell>{i18n.t("No pages found")}</TableCell>
              </TableRow>
            ),
          )}
        </TableBody>
      </Table>
    </Card>
  ),
);
export default DirectoryRootList;
