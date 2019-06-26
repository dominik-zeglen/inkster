import * as React from "react";
import { Plus } from "react-feather";
import IconButton from "aurora-ui-kit/dist/components/IconButton";

import { ListViewProps } from "../../";
import i18n from "../../i18n";
import Container from "../../components/Container";
import PageHeader from "../../components/PageHeader";
import DirectoryRootList from "./DirectoryRootList";
import { RootDirectories_getRootDirectories_edges_node } from "../queries/types/RootDirectories";
import createUseStyles from "aurora-ui-kit/dist/utils/jss";
import { ITheme } from "aurora-ui-kit/dist/theme";

interface Props extends ListViewProps {
  directories: RootDirectories_getRootDirectories_edges_node[];
}

const useStyles = createUseStyles((theme: ITheme) => ({
  root: {
    display: "grid" as "grid",
    gridColumnGap: theme.spacing,
    gridTemplateColumns: "2fr 1fr",
  },
}));
export const DirectoryRootPage: React.FC<Props> = ({
  directories,
  disabled,
  loading,
  pageInfo,
  onAdd,
  onNextPage,
  onPreviousPage,
  onRowClick,
}) => {
  const classes = useStyles();
  return (
    <Container width="md">
      <PageHeader title={i18n.t("Directories")}>
        <IconButton disabled={disabled || loading} onClick={onAdd}>
          <Plus />
        </IconButton>
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
  );
};
export default DirectoryRootPage;
