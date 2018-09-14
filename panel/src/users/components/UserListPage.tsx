import * as React from "react";
import { Plus } from "react-feather";
import withStyles from "react-jss";

import ActionDialog from "../../components/ActionDialog";
import Container from "../../components/Container";
import Form from "../../components/Form";
import i18n from "../../i18n";
import IconButton from "../../components/IconButton";
import Input from "../../components/Input";
import { ListViewProps } from "../../";
import PageHeader from "../../components/PageHeader";
import Toggle from "../../components/Toggle";
import UserList from "./UserList";

interface Props extends ListViewProps<{ email: string }> {
  users?: Array<{
    id: string;
    email: string;
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
  { displayName: "UserRootPage" }
);
export const UserListPage = decorate<Props>(
  ({
    classes,
    disabled,
    loading,
    pageInfo,
    users,
    onAdd,
    onNextPage,
    onPreviousPage,
    onRowClick
  }) => (
    <Toggle>
      {(openedAddUserDialog, { toggle: toggleAddUserDialog }) => (
        <>
          <Container width="md">
            <PageHeader title={i18n.t("Users")}>
              <IconButton
                disabled={disabled || loading}
                icon={Plus}
                onClick={toggleAddUserDialog}
              />
            </PageHeader>
            <div className={classes.root}>
              <div>
                <UserList
                  disabled={disabled}
                  loading={loading}
                  users={users}
                  pageInfo={pageInfo}
                  onNextPage={onNextPage}
                  onPreviousPage={onPreviousPage}
                  onRowClick={onRowClick}
                />
              </div>
              <div />
            </div>
          </Container>
          <Form initial={{ email: "" }} onSubmit={onAdd}>
            {({ change, data, submit }) => (
              <ActionDialog
                show={openedAddUserDialog}
                size="xs"
                onClose={toggleAddUserDialog}
                onConfirm={submit}
                title={i18n.t("Add new user")}
              >
                <Input
                  name="email"
                  onChange={change}
                  value={data.email}
                  label={i18n.t("User email")}
                  type="email"
                />
              </ActionDialog>
            )}
          </Form>
        </>
      )}
    </Toggle>
  )
);
export default UserListPage;
