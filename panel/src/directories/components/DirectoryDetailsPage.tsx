import * as React from "react";
import withStyles from "react-jss";
import { Trash } from "react-feather";

import Container from "../../components/Container";
import Form from "../../components/Form";
import FormSave from "../../components/FormSave";
import PageHeader from "../../components/PageHeader";
import IconButton from "../../components/IconButton";
import { FormViewProps, ListViewProps } from "../../";
import DirectoryProperties from "./DirectoryProperties";
import DirectoryPages from "./DirectoryPages";
import DirectoryStatus from "./DirectoryStatus";
import { Directory_getDirectory } from "../queries/types/Directory";
import { maybe } from "../../utils";

interface FormData {
  name: string;
  isPublished: boolean;
}
interface Props extends FormViewProps<FormData>, ListViewProps<{}> {
  directory: Directory_getDirectory;
  onDelete: () => void;
}

const decorate = withStyles(
  (theme: any) => ({
    root: {
      display: "grid" as "grid",
      gridColumnGap: theme.spacing + "px",
      gridTemplateColumns: "2fr 1fr",
    },
  }),
  { displayName: "DirectoryDetailsPage" },
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
    onRowClick,
  }) => (
    <Form
      initial={{
        isPublished:
          directory && directory.isPublished ? directory.isPublished : false,
        name: directory && directory.name ? directory.name : "",
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
              onClick={onDelete}
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
                pages={maybe(() =>
                  directory.pages.edges.map(edge => edge.node),
                )}
                disabled={disabled || loading}
                pageInfo={pageInfo}
                onAdd={onAdd}
                onNextPage={onNextPage}
                onPreviousPage={onPreviousPage}
                onRowClick={onRowClick}
              />
            </div>
            <div>
              <DirectoryStatus
                createdAt={directory ? directory.createdAt : undefined}
                data={data}
                updatedAt={directory ? directory.updatedAt : undefined}
                onChange={change}
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
  ),
);
export default DirectoryDetailsPage;
