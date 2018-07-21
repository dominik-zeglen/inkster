import * as React from "react";
import withStyles from "react-jss";
import { Trash } from "react-feather";

import ActionDialog from "../../components/ActionDialog";
import Container from "../../components/Container";
import Form from "../../components/Form";
import FormSave from "../../components/FormSave";
import Toggle from "../../components/Toggle";
import PageHeader from "../../components/PageHeader";
import IconButton from "../../components/IconButton";
import { FormViewProps, ListViewProps } from "../../";
import DirectoryProperties from "./DirectoryProperties";
import DirectoryPages from "./DirectoryPages";
import i18n from "../../i18n";

interface FormData {
  name: string;
}
interface Props extends FormViewProps<FormData>, ListViewProps<{}> {
  directory?: {
    id: string;
    name?: string;
    pages?: Array<{
      id: string;
      name?: string;
    }>;
  };
  onDelete: () => void;
}

const decorate = withStyles(
  (theme: any) => ({
    root: {
      display: "grid" as "grid",
      gridColumnGap: theme.spacing + "px",
      gridTemplateColumns: "2fr 1fr"
    }
  }),
  { displayName: "DirectoryDetailsPage" }
);
export const DirectoryDetailsPage = decorate<Props>(
  ({
    classes,
    directory,
    disabled,
    loading,
    transaction,
    pageInfo,
    onAdd,
    onBack,
    onDelete,
    onSubmit,
    onNextPage,
    onPreviousPage,
    onRowClick
  }) => (
    <Toggle>
      {(openedDeleteDialog, { toggle: toggleDeleteDialog }) => (
        <>
          <Form
            initial={{
              name: directory && directory.name ? directory.name : ""
            }}
            onSubmit={onSubmit}
            key={JSON.stringify(directory)}
          >
            {({ change, data, hasChanged, submit }) => (
              <Container width="md">
                <PageHeader
                  title={directory ? directory.name : undefined}
                  onBack={onBack}
                >
                  <IconButton
                    disabled={disabled || loading}
                    icon={Trash}
                    onClick={toggleDeleteDialog}
                  />
                </PageHeader>
                <div className={classes.root}>
                  <div>
                    <DirectoryProperties
                      data={data}
                      disabled={disabled || loading}
                      onChange={change}
                    />
                    <DirectoryPages
                      pages={directory ? directory.pages : undefined}
                      disabled={disabled || loading}
                      pageInfo={pageInfo}
                      onAdd={onAdd}
                      onNextPage={onNextPage}
                      onPreviousPage={onPreviousPage}
                      onRowClick={onRowClick}
                    />
                  </div>
                </div>
                <FormSave
                  disabled={disabled || !hasChanged}
                  variant={transaction}
                  onConfirm={submit}
                />
              </Container>
            )}
          </Form>
          {!(disabled || loading) &&
            directory &&
            directory.name && (
              <ActionDialog
                show={openedDeleteDialog}
                size="xs"
                title={i18n.t("Remove directory")}
                onClose={toggleDeleteDialog}
                onConfirm={onDelete}
              >
                {i18n.t("Are you sure you want to remove {{ name }}?", {
                  name: directory.name
                })}
              </ActionDialog>
            )}
        </>
      )}
    </Toggle>
  )
);
export default DirectoryDetailsPage;
