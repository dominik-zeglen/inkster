import * as React from "react";
import withStyles from "react-jss";
import { Button, Panel } from "react-bootstrap";
import { AlertTriangle } from "react-feather";

import Checkbox from "../../components/Checkbox";
import Container from "../../components/Container";
import Form from "../../components/Form";
import Input from "../../components/Input";
import i18n from "../../i18n";

export interface FormData {
  email: string;
  password: string;
  remember: boolean;
}

export interface Props {
  disabled: boolean;
  error: boolean;
  onPasswordRecovery: () => void;
  onSubmit: (data: FormData) => void;
}

const initialForm: FormData = {
  email: "",
  password: "",
  remember: false
};
const decorate = withStyles(theme => ({
  buttonContainer: {
    display: "flex" as "flex",
    justifyContent: "flex-end" as "flex-end"
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
      width: 36
    },
    display: "grid" as "grid",
    gridColumnGap: theme.spacing + "px",
    gridTemplateColumns: "36px 1fr",
  },
  forgotPasswordLink: {
    cursor: "pointer" as "pointer",
    fontWeight: 600 as 600
  },
  panel: {
    padding: 60,
    width: "100%"
  },
  root: {
    alignItems: "center" as "center",
    display: "flex" as "flex",
    height: "100vh",
    padding: 80
  }
}));
export const LoginPage = decorate<Props>(
  ({ classes, disabled, error, onPasswordRecovery, onSubmit }) => (
    <Form initial={initialForm} onSubmit={onSubmit}>
      {({ change, data, hasChanged }) => (
        <Container width="xs">
          <div className={classes.root}>
            <Panel className={classes.panel}>
              {error && (
                <Panel className={classes.errorPanel}>
                  <Panel.Body>
                    <div className={classes.errorPanelContent}>
                      <AlertTriangle />
                      <div>
                        <div>{i18n.t("Username or password is invalid.")}</div>
                        <div className={classes.forgotPasswordLink}>
                          {i18n.t("Forgot your password?")}
                        </div>
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
                    {i18n.t("Login", { context: "button" })}
                  </Button>
                </div>
              </div>
            </Panel>
          </div>
        </Container>
      )}
    </Form>
  )
);
export default LoginPage;
