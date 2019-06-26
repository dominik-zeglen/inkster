import * as React from "react";
import { Plus } from "react-feather";
import IconButton from "aurora-ui-kit/dist/components/IconButton";

import Container from "../../components/Container";
import i18n from "../../i18n";
import { ListViewProps } from "../../";
import PageHeader from "../../components/PageHeader";
import UserList from "./UserList";
import { UserList_users_edges_node } from "../queries/types/UserList";
import createUseStyles from "aurora-ui-kit/dist/utils/jss";
import { ITheme } from "aurora-ui-kit/dist/theme";

interface Props extends ListViewProps {
  users: UserList_users_edges_node[];
}

const useStyles = createUseStyles((theme: ITheme) => ({
  root: {
    display: "grid" as "grid",
    gridColumnGap: theme.spacing,
    gridTemplateColumns: "2fr 1fr",
  },
}));
export const UserListPage: React.FC<Props> = ({
  disabled,
  loading,
  pageInfo,
  users,
  onAdd,
  onNextPage,
  onPreviousPage,
  onRowClick,
}) => {
  const classes = useStyles();
  return (
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
  );
};
export default UserListPage;
