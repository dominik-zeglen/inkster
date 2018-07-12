import * as React from "react";
import { Plus } from "react-feather";
import withStyles from "react-jss";

import { ListViewProps } from "../../";
import i18n from "../../i18n";
import ActionDialog from "../../components/ActionDialog";
import Container from "../../components/Container";
import Form from "../../components/Form";
import Input from "../../components/Input";
import PageHeader from "../../components/PageHeader";
import IconButton from "../../components/IconButton";
import Toggle from "../../components/Toggle";
import DirectoryRootList from "./DirectoryRootList";

interface Props extends ListViewProps<{name: string}> {
  directories?: Array<{
    id: string;
    name: string;
  }>;
}

const decorate = withStyles(
  (theme: any) => ({
    root: {
      display: "grid" as "grid",
      gridColumnGap: theme.spacing,
      gridTemplateColumns: "2fr 1fr"
    }
  }),
  { displayName: "DirectoryRootPage" }
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
    onRowClick
  }) => (
    <Toggle>
      {(openedAddDirectoryDialog, { toggle: toggleAddDirectoryDialog }) => (
        <>
          <Container width="md">
            <PageHeader title={i18n.t("Directories")}>
              <IconButton
                disabled={disabled || loading}
                icon={Plus}
                onClick={toggleAddDirectoryDialog}
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
          <Form initial={{ name: '' }} onSubmit={onAdd}>
            {({ change, data, submit }) => (
              <ActionDialog
                show={openedAddDirectoryDialog}
                size="xs"
                onClose={toggleAddDirectoryDialog}
                onConfirm={submit}
                title={i18n.t("Add new directory")}
              >
                <Input
                  name="name"
                  onChange={change}
                  value={data.name}
                  label={i18n.t("Directory name")}
                />
              </ActionDialog>
            )}
          </Form>
        </>
      )}
    </Toggle>
  )
);
export default DirectoryRootPage;
