import * as React from "react";
import { Plus, Trash } from "react-feather";
import withStyles from "react-jss";

import PageHeader from "../../components/PageHeader";
import Container from "../../components/Container";
import IconButton from "../../components/IconButton";
import ActionDialog from "../../components/ActionDialog";
import Input from "../../components/Input";
import Form from "../../components/Form";
import FormSave from "../../components/FormSave";
import Toggle from "../../components/Toggle";
import { ViewProps, FormViewProps } from "../../";
import i18n from "../../i18n";
import PageProperties from "./PageProperties";
import PageFieldProperties from "./PageFieldProperties";

interface PageField {
  id: string;
  name: string;
  type: string;
  value: string;
}
export interface FormData {
  name: string;
  slug: string;
  fields: PageField[];
  addFields: PageField[];
  removeFields: string[];
}
interface Props extends ViewProps, FormViewProps<FormData> {
  page?: {
    id: string;
    name?: string;
    slug?: string;
    parent?: {
      id: string;
      name?: string;
    };
    fields: PageField[];
  };
}

const decorate = withStyles(
  (theme: any) => ({
    root: {
      display: "grid" as "grid",
      gridColumnGap: theme.spacing + "px",
      gridTemplateColumns: "2fr 1fr"
    }
  }),
  { displayName: "PageDetailsPage" }
);
export const PageDetailsPage = decorate<Props>(
  ({
    classes,
    disabled,
    loading,
    page,
    transaction,
    onBack,
    onDelete,
    onSubmit
  }) => {
    const initialForm = {
      name: page && page.name ? page.name : "",
      slug: page && page.slug ? page.slug : "",
      fields:
        page && page.fields
          ? (page.fields.map(f => ({ id: f.name, ...f })) as PageField[])
          : [],
      addFields: [] as PageField[],
      removeFields: [] as string[]
    };
    return (
      <Form
        initial={initialForm}
        onSubmit={onSubmit}
        key={JSON.stringify(page ? JSON.stringify(page) : "loading")}
      >
        {({ change, data, hasChanged, submit }) => {
          const handleFieldAdd = (field: { name: string; type: string }) =>
            change({
              target: {
                name: "addFields",
                value: [
                  ...data.addFields,
                  { ...field, id: "new-" + data.addFields.length }
                ]
              }
            } as any);
          const handleFieldRemove = (name: string, id: string) => () => {
            change({
              target: {
                name,
                value: data[name].filter((f: PageField) => f.id !== id)
              }
            } as any);
            if (name === "fields") {
              change({
                target: {
                  name: "removeFields",
                  value: [id, ...data.removeFields]
                }
              } as any);
            }
          };
          const handleChange = (name: string, id: string) => (
            event: React.ChangeEvent<any>
          ) =>
            change({
              target: {
                name,
                value: data[name].map(
                  (f: PageField) =>
                    f.id === id
                      ? { ...f, [event.target.name]: event.target.value }
                      : f
                )
              }
            } as any);
          return (
            <Toggle>
              {(openedRemoveDialog, { toggle: toggleRemoveDialog }) => (
                <Toggle>
                  {(openedFieldAddDialog, { toggle: toggleFieldAddDialog }) => (
                    <Toggle>
                      {(
                        openedFieldRemoveDialog,
                        { toggle: toggleFieldRemoveDialog }
                      ) => (
                        <>
                          <Container width="md">
                            <PageHeader
                              onBack={onBack}
                              title={page ? page.name : undefined}
                            >
                              <IconButton
                                disabled={disabled || loading}
                                icon={Plus}
                                onClick={toggleFieldAddDialog}
                              />
                              <IconButton
                                disabled={disabled || loading}
                                icon={Trash}
                                onClick={toggleRemoveDialog}
                              />
                            </PageHeader>
                            <div className={classes.root}>
                              <div>
                                <PageProperties
                                  data={data}
                                  disabled={disabled || loading}
                                  onChange={change}
                                />
                                {data.fields.map((field, index) => (
                                  <PageFieldProperties
                                    data={field}
                                    key={field.id + index}
                                    name="fields"
                                    onChange={handleChange("fields", field.id)}
                                    onDelete={handleFieldRemove(
                                      "fields",
                                      field.id
                                    )}
                                  />
                                ))}
                                {data.addFields.map((field, index) => (
                                  <PageFieldProperties
                                    data={field}
                                    key={field.id + index}
                                    name="addFields"
                                    onChange={handleChange(
                                      "addFields",
                                      field.id
                                    )}
                                    onDelete={handleFieldRemove(
                                      "addFields",
                                      field.id
                                    )}
                                  />
                                ))}
                              </div>
                            </div>
                            <FormSave
                              disabled={disabled || loading || !hasChanged}
                              onConfirm={submit}
                              variant={transaction}
                            />
                          </Container>
                          {!disabled &&
                            !loading &&
                            page && (
                              <>
                                <ActionDialog
                                  show={openedRemoveDialog}
                                  size="xs"
                                  title={i18n.t("Remove page")}
                                  onClose={toggleRemoveDialog}
                                  onConfirm={onDelete}
                                >
                                  {i18n.t(
                                    "Are you sure you want to remove {{ name }}?",
                                    { name: page.name }
                                  )}
                                </ActionDialog>
                                <Form
                                  initial={{ type: "", name: "" }}
                                  onSubmit={handleFieldAdd}
                                >
                                  {({
                                    change: handleAddFieldChange,
                                    data: addFieldData,
                                    submit: addField
                                  }) => (
                                    <ActionDialog
                                      show={openedFieldAddDialog}
                                      size="xs"
                                      title={i18n.t("Add page field")}
                                      onClose={toggleFieldAddDialog}
                                      onConfirm={addField as () => void}
                                    >
                                      <Input
                                        name="name"
                                        label={i18n.t("Name")}
                                        value={addFieldData.name}
                                        onChange={handleAddFieldChange}
                                      />
                                      <Input
                                        name="type"
                                        label={i18n.t("Type")}
                                        value={addFieldData.type}
                                        onChange={handleAddFieldChange}
                                        type="select"
                                      >
                                        <>
                                          <option value="text">
                                            {i18n.t("Short text")}
                                          </option>
                                          <option
                                            value="longText"
                                            selected={true}
                                          >
                                            {i18n.t("Long text")}
                                          </option>
                                        </>
                                      </Input>
                                    </ActionDialog>
                                  )}
                                </Form>
                              </>
                            )}
                        </>
                      )}
                    </Toggle>
                  )}
                </Toggle>
              )}
            </Toggle>
          );
        }}
      </Form>
    );
  }
);
export default PageDetailsPage;
