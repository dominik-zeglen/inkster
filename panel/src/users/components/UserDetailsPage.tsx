import * as React from "react";
import withStyles from "react-jss";
import { Trash } from "react-feather";

import ActionDialog from "../../components/ActionDialog";
import Container from "../../components/Container";
import Form from "../../components/Form";
import FormSave from "../../components/FormSave";
import { FormViewProps } from "../../";
import i18n from "../../i18n";
import IconButton from "../../components/IconButton";
import PageHeader from "../../components/PageHeader";
import Toggle from "../../components/Toggle";
import UserStatus from './UserStatus'
import UserProperties from './UserProperties'

interface FormData {
  email: string;
  isActive: boolean;
}
interface Props extends FormViewProps<FormData> {
  user?: {
    id: string;
    email: string;
    isActive: boolean;
    createdAt: string;
    updatedAt: string;
  };
  onDelete: () => void;
}

const decorate = withStyles(
  theme => ({
    cardContainer: {
      marginBottom: theme.spacing
    },
    root: {
      display: "grid" as "grid",
      gridColumnGap: theme.spacing + "px",
      gridTemplateColumns: "2fr 1fr"
    }
  }),
  { displayName: "UserDetailsPage" }
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
    onSubmit
  }) => (
    <Toggle>
      {(openedDeleteDialog, { toggle: toggleDeleteDialog }) => (
        <>
          <Form
            initial={{
              email: user ? user.email : "",
              isActive: user ? user.isActive : false
            }}
            onSubmit={onSubmit}
            key={JSON.stringify(user)}
          >
            {({ change, data, hasChanged, submit }) => (
              <Container width="md">
                <PageHeader
                  title={user ? user.email : undefined}
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
                    <UserProperties data={data} disabled={disabled || loading} onChange={change} />
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
          <ActionDialog
            show={openedDeleteDialog}
            size="xs"
            title={i18n.t("Remove user")}
            onClose={toggleDeleteDialog}
            onConfirm={onDelete}
          >
            {i18n.t("Are you sure you want to remove {{ email }}?", {
              email: user ? user.email : ""
            })}
          </ActionDialog>
        </>
      )}
    </Toggle>
  )
);
export default UserDetailsPage;
