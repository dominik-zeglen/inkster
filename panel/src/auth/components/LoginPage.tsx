import * as React from "react";
import withStyles from "react-jss";
import { AlertTriangle } from "react-feather";
import Checkbox from "aurora-ui-kit/dist/components/Checkbox";
import Button from "aurora-ui-kit/dist/components/Button";
import Card from "aurora-ui-kit/dist/components/Card";
import CardContent from "aurora-ui-kit/dist/components/CardContent";
import Input from "aurora-ui-kit/dist/components/TextInput";

import PageLayout from "./PageLayout";
import Form from "../../components/Form";
import Link from "../../components/Link";
import Typography from "../../components/Typography";
import i18n from "../../i18n";
import Spacer from "../../components/Spacer";

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
  ({ classes, error, passwordRecoveryHref, onSubmit }) => (
    <Form initial={initialForm} onSubmit={onSubmit}>
      {({ change, data, submit }) => (
        <PageLayout
          header={i18n.t("Log in", {
            context: "header",
          })}
        >
          {error && (
            <Card className={classes.errorPanel}>
              <CardContent>
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
              </CardContent>
            </Card>
          )}
          <div>
            <Input
              label={i18n.t("E-mail")}
              InputProps={{
                componentProps: {
                  type: "email",
                },
              }}
              value={data.email}
              onChange={value =>
                change({
                  target: {
                    name: "email",
                    value,
                  },
                } as any)
              }
            />
            <Spacer />
            <Input
              label={i18n.t("Password")}
              InputProps={{
                componentProps: {
                  type: "password",
                },
              }}
              value={data.password}
              onChange={value =>
                change({
                  target: {
                    name: "password",
                    value,
                  },
                } as any)
              }
            />
            <Checkbox
              checked={data.remember}
              label={i18n.t("Remember me")}
              onChange={() =>
                change({
                  target: {
                    name: "remember",
                    value: !data.remember,
                  },
                } as any)
              }
            />
            <div className={classes.buttonContainer}>
              <Button color="primary" type="submit" onClick={submit}>
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
