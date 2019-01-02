import * as React from "react";
import { Plus } from "react-feather";
import withStyles from "react-jss";

import { ListViewProps } from "../../";
import i18n from "../../i18n";
import Container from "../../components/Container";
import PageHeader from "../../components/PageHeader";
import IconButton from "../../components/IconButton";
import DirectoryRootList from "./DirectoryRootList";
import { RootDirectories_getRootDirectories_edges_node } from "../queries/types/RootDirectories";

interface Props extends ListViewProps<{ name: string }> {
  directories: RootDirectories_getRootDirectories_edges_node[];
}

const decorate = withStyles(
  (theme: any) => ({
    root: {
      display: "grid" as "grid",
      gridColumnGap: theme.spacing,
      gridTemplateColumns: "2fr 1fr",
    },
  }),
  { displayName: "DirectoryRootPage" },
);
export const DirectoryRootPage = decorate<Props>(
  ({
    classes,
    directories,
    disabled,
    loading,
    pageInfo,
    onAdd,
    onNextPage,
    onPreviousPage,
    onRowClick,
  }) => (
    <Container width="md">
      <PageHeader title={i18n.t("Directories")}>
        <IconButton
          disabled={disabled || loading}
          icon={Plus}
          onClick={onAdd}
        />
      </PageHeader>
      <div className={classes.root}>
        <div>
          <DirectoryRootList
            disabled={disabled}
            loading={loading}
            directories={directories}
            pageInfo={pageInfo}
            onNextPage={onNextPage}
            onPreviousPage={onPreviousPage}
            onRowClick={onRowClick}
          />
        </div>
        <div />
      </div>
    </Container>
  ),
);
export default DirectoryRootPage;
