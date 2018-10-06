import * as React from "react";
import withStyles from "react-jss";
import { Button, Panel } from "react-bootstrap";
import { AlertTriangle } from "react-feather";

import PageLayout from "./PageLayout";
import Checkbox from "../../components/Checkbox";
import Form from "../../components/Form";
import Link from "../../components/Link";
import Input from "../../components/Input";
import Typography from "../../components/Typography";
import i18n from "../../i18n";

export interface FormData {
  email: string;
  password: string;
  remember: boolean;
}

export interface Props {
  disabled: boolean;
  error: boolean;
  passwordRecoveryHref: string;
  onSubmit: (data: FormData) => void;
}

const initialForm: FormData = {
  email: "",
  password: "",
  remember: false,
};
const decorate = withStyles(theme => ({
  buttonContainer: {
    display: "flex" as "flex",
    justifyContent: "flex-end" as "flex-end",
  },
  errorPanel: {
    backgroundColor: theme.colors.error.main,
    color: theme.colors.white.main,
    fontSize: theme.typography.caption.fontSize,
    marginBottom: theme.spacing * 4,
  },
  errorPanelContent: {
    "& svg": {
      height: 36,
      width: 36,
    },
    display: "grid" as "grid",
    gridColumnGap: theme.spacing + "px",
    gridTemplateColumns: "36px 1fr",
  },
  errorPanelForgotPasswordLink: {
    cursor: "pointer" as "pointer",
    fontWeight: 600 as 600,
  },
  forgotPasswordLink: {
    display: "block" as "block",
    marginTop: theme.spacing * 2,
    textAlign: "center" as "center",
  },
}));
export const LoginPage = decorate<Props>(
  ({ classes, disabled, error, passwordRecoveryHref, onSubmit }) => (
    <Form initial={initialForm} onSubmit={onSubmit}>
      {({ change, data, hasChanged }) => (
        <PageLayout
          header={i18n.t("Log in", {
            context: "header",
          })}
        >
          {error && (
            <Panel className={classes.errorPanel}>
              <Panel.Body>
                <div className={classes.errorPanelContent}>
                  <AlertTriangle />
                  <div>
                    <div>{i18n.t("Username or password is invalid.")}</div>
                    <Link href={passwordRecoveryHref}>
                      <div className={classes.errorPanelForgotPasswordLink}>
                        {i18n.t("Forgot your password?")}
                      </div>
                    </Link>
                  </div>
                </div>
              </Panel.Body>
            </Panel>
          )}
          <div>
            <Input
              label={i18n.t("E-mail")}
              name="email"
              type="email"
              value={data.email}
              onChange={change}
            />
            <Input
              label={i18n.t("Password")}
              name="password"
              type="password"
              value={data.password}
              onChange={change}
            />
            <Checkbox
              label={i18n.t("Remember me")}
              name="remember"
              value={data.remember}
              onChange={change}
            />
            <div className={classes.buttonContainer}>
              <Button bsStyle="primary" type="submit">
                {i18n.t("Log in", { context: "button" })}
              </Button>
            </div>
            <Link
              className={classes.forgotPasswordLink}
              href={passwordRecoveryHref}
            >
              <Typography component="span" variant="anchor">
                {i18n.t("Reset password", {
                  context: "link",
                })}
              </Typography>
            </Link>
          </div>
        </PageLayout>
      )}
    </Form>
  ),
);
export default LoginPage;
