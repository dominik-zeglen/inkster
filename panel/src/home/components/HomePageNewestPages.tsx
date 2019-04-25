import * as React from "react";
import Card from "aurora-ui-kit/dist/components/Card";
import CardHeader from "aurora-ui-kit/dist/components/CardHeader";
import CardTitle from "aurora-ui-kit/dist/components/CardTitle";
import TableHead from "aurora-ui-kit/dist/components/TableHead";
import TableCell from "aurora-ui-kit/dist/components/TableCell";
import Table from "aurora-ui-kit/dist/components/Table";
import TableBody from "aurora-ui-kit/dist/components/TableBody";
import TableRow from "aurora-ui-kit/dist/components/TableRow";
import Skeleton from "aurora-ui-kit/dist/components/Skeleton";
import withStyles, { CSSProperties } from "react-jss";

import i18n from "../../i18n";
import { renderCollection, maybe } from "../../utils";
import { Viewer_viewer_pages_edges_node } from "../queries/types/Viewer";

interface Props {
  disabled: boolean;
  pages: Viewer_viewer_pages_edges_node[];
  onPageClick: (id: string) => void;
}

const styles: CSSProperties = {
  row: {
    cursor: "pointer",
  },
};
export const HomePageNewestPages = withStyles(styles)<Props>(
  ({ classes, disabled, pages, onPageClick }) => (
    <Card>
      <CardHeader>
        <CardTitle>{i18n.t("Your newest pages")}</CardTitle>
      </CardHeader>
      <Table>
        <TableHead>
          <TableCell>{i18n.t("Page Title")}</TableCell>
        </TableHead>
        <TableBody>
          {renderCollection(pages, page => (
            <TableRow
              className={classes.row}
              hover={!disabled}
              onClick={page ? () => onPageClick(page.id) : undefined}
            >
              <TableCell>
                {maybe<React.ReactNode>(() => page.name, <Skeleton />)}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </Card>
  ),
);
export default HomePageNewestPages;
