import * as React from "react";
import withStyles from "react-jss";
import { Button } from "react-bootstrap";

import PageLayout from "./PageLayout";
import Form from "../../components/Form";
import Input from "../../components/Input";
import i18n from "../../i18n";

export interface FormData {
  password: string;
  passwordConfirm: string;
}

export interface Props {
  disabled: boolean;
  onSubmit: (data: FormData) => void;
}

const initialForm: FormData = {
  password: "",
  passwordConfirm: ""
};
const decorate = withStyles(theme => ({
  buttonContainer: {
    display: "flex" as "flex",
    justifyContent: "flex-end" as "flex-end"
  }
}));
export const PasswordResetPage = decorate<Props>(
  ({ classes, disabled, onSubmit }) => (
    <Form initial={initialForm} onSubmit={onSubmit}>
      {({ change, data, hasChanged }) => (
        <PageLayout
          header={i18n.t("Reset password", {
            context: "header"
          })}
        >
          <Input
            label={i18n.t("New password", {
              context: "label"
            })}
            name="password"
            type="password"
            value={data.password}
            onChange={change}
          />
          <Input
            error={data.password !== data.passwordConfirm}
            label={i18n.t("Confirm password", {
              context: "label"
            })}
            helperText={
              data.password !== data.passwordConfirm
                ? i18n.t("Passwords do not match", {
                    context: "caption"
                  })
                : undefined
            }
            name="passwordConfirm"
            type="password"
            value={data.passwordConfirm}
            onChange={change}
          />
          <div className={classes.buttonContainer}>
            <Button
              bsStyle="primary"
              disabled={
                disabled || !hasChanged || data.password !== data.passwordConfirm
              }
              type="submit"
            >
              {i18n.t("Submit", { context: "button" })}
            </Button>
          </div>
        </PageLayout>
      )}
    </Form>
  )
);
export default PasswordResetPage;
