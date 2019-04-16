import * as React from "react";
import { Plus } from "react-feather";
import withStyles from "react-jss";
import IconButton from "aurora-ui-kit/dist/components/IconButton";

import Container from "../../components/Container";
import i18n from "../../i18n";
import { ListViewProps } from "../../";
import PageHeader from "../../components/PageHeader";
import UserList from "./UserList";
import { UserList_users_edges_node } from "../queries/types/UserList";

interface Props extends ListViewProps {
  users: UserList_users_edges_node[];
}

const decorate = withStyles(
  (theme: any) => ({
    root: {
      display: "grid" as "grid",
      gridColumnGap: theme.spacing,
      gridTemplateColumns: "2fr 1fr",
    },
  }),
  { displayName: "UserRootPage" },
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
    onRowClick,
  }) => (
    <Container width="md">
      <PageHeader title={i18n.t("Users")}>
        <IconButton disabled={disabled || loading} onClick={onAdd}>
          <Plus />
        </IconButton>
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
  ),
);
export default UserListPage;
