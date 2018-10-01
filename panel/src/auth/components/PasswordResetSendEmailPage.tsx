import * as React from "react";
import withStyles from "react-jss";
import { Button } from "react-bootstrap";

import PageLayout from "./PageLayout";
import Form from "../../components/Form";
import Input from "../../components/Input";
import i18n from "../../i18n";

export interface FormData {
  email: string;
}

export interface Props {
  disabled: boolean;
  onSubmit: (data: FormData) => void;
}

const initialForm: FormData = {
  email: ""
};
const decorate = withStyles(theme => ({
  buttonContainer: {
    display: "flex" as "flex",
    justifyContent: "flex-end" as "flex-end"
  }
}));
export const PasswordResetSendEmailPage = decorate<Props>(
  ({ classes, disabled, onSubmit }) => (
    <Form initial={initialForm} onSubmit={onSubmit}>
      {({ change, data, hasChanged }) => (
        <PageLayout
          header={i18n.t("Reset password", {
            context: "header"
          })}
        >
          <Input
            label={i18n.t("E-mail")}
            name="email"
            helperText={i18n.t(
              "Enter your e-mail address, so we can send you message with further instructions",
              {
                context: "caption"
              }
            )}
            type="email"
            value={data.email}
            onChange={change}
          />
          <div className={classes.buttonContainer}>
            <Button
              bsStyle="primary"
              disabled={disabled || !hasChanged}
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
export default PasswordResetSendEmailPage;
