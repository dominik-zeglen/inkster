import * as React from "react";
import { Trash } from "react-feather";
import IconButton from "aurora-ui-kit/dist/components/IconButton";

import Container from "../../components/Container";
import Form from "../../components/Form";
import FormSave from "../../components/FormSave";
import { FormViewProps, PaginatedListProps } from "../../";
import PageHeader from "../../components/PageHeader";
import UserStatus from "./UserStatus";
import UserProperties from "./UserProperties";
import { UserDetails_user } from "../queries/types/UserDetails";
import Spacer from "../../components/Spacer";
import UserPages from "./UserPages";
import { maybe } from "../../utils";
import createUseStyles from "aurora-ui-kit/dist/utils/jss";
import { ITheme } from "aurora-ui-kit/dist/theme";

interface FormData {
  email: string;
  isActive: boolean;
}
interface Props extends FormViewProps<FormData>, PaginatedListProps {
  user: UserDetails_user;
  onDelete: () => void;
}

const useStyles = createUseStyles((theme: ITheme) => ({
  cardContainer: {
    marginBottom: theme.spacing,
  },
  root: {
    display: "grid" as "grid",
    gridColumnGap: theme.spacing + "px",
    gridTemplateColumns: "2fr 1fr",
  },
}));
export const UserDetailsPage: React.FC<Props> = ({
  disabled,
  loading,
  transaction,
  user,
  onBack,
  onDelete,
  onSubmit,
  ...listProps
}) => {
  const classes = useStyles();
  return (
    <Form
      initial={{
        email: user ? user.email : "",
        isActive: user ? user.isActive : false,
      }}
      onSubmit={onSubmit}
      key={JSON.stringify(user)}
    >
      {({ change, data, hasChanged }) => (
        <Container width="md">
          <PageHeader title={user ? user.email : undefined} onBack={onBack}>
            <IconButton disabled={disabled || loading} onClick={onDelete}>
              <Trash />
            </IconButton>
          </PageHeader>
          <div className={classes.root}>
            <div>
              <UserProperties
                data={data}
                disabled={disabled || loading}
                onChange={change}
              />
              <Spacer />
              <UserPages
                disabled={disabled}
                loading={loading}
                pages={maybe(() => user.pages.edges.map(edge => edge.node))}
                {...listProps}
              />
            </div>
            <div>
              <UserStatus
                createdAt={user ? user.createdAt : undefined}
                data={data}
                updatedAt={user ? user.updatedAt : undefined}
                onChange={change}
              />
            </div>
          </div>
          <FormSave
            disabled={disabled || !hasChanged}
            variant={transaction}
            onConfirm={() => undefined}
          />
        </Container>
      )}
    </Form>
  );
};
export default UserDetailsPage;
