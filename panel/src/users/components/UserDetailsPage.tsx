import * as React from "react";
import withStyles from "react-jss";
import { Trash } from "react-feather";

import Container from "../../components/Container";
import Form from "../../components/Form";
import FormSave from "../../components/FormSave";
import { FormViewProps } from "../../";
import IconButton from "../../components/IconButton";
import PageHeader from "../../components/PageHeader";
import UserStatus from "./UserStatus";
import UserProperties from "./UserProperties";
import { UserDetails_user } from "../queries/types/UserDetails";

interface FormData {
  email: string;
  isActive: boolean;
}
interface Props extends FormViewProps<FormData> {
  user: UserDetails_user;
  onDelete: () => void;
}

const decorate = withStyles(
  theme => ({
    cardContainer: {
      marginBottom: theme.spacing,
    },
    root: {
      display: "grid" as "grid",
      gridColumnGap: theme.spacing + "px",
      gridTemplateColumns: "2fr 1fr",
    },
  }),
  { displayName: "UserDetailsPage" },
);
export const UserDetailsPage = decorate<Props>(
  ({
    classes,
    disabled,
    loading,
    transaction,
    user,
    onBack,
    onDelete,
    onSubmit,
  }) => (
    <Form
      initial={{
        email: user ? user.email : "",
        isActive: user ? user.isActive : false,
      }}
      onSubmit={onSubmit}
      key={JSON.stringify(user)}
    >
      {({ change, data, hasChanged, submit }) => (
        <Container width="md">
          <PageHeader title={user ? user.email : undefined} onBack={onBack}>
            <IconButton
              disabled={disabled || loading}
              icon={Trash}
              onClick={onDelete}
            />
          </PageHeader>
          <div className={classes.root}>
            <div>
              <UserProperties
                data={data}
                disabled={disabled || loading}
                onChange={change}
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
            onConfirm={submit}
          />
        </Container>
      )}
    </Form>
  ),
);
export default UserDetailsPage;
