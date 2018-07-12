import * as React from "react";
import { Button } from "react-bootstrap";
import withStyles from "react-jss";

import { StandardProps } from "./";
import { TransactionState } from "../";
import i18n from "../i18n";

interface Props extends StandardProps {
  disabled: boolean;
  variant: TransactionState;
  onConfirm: (event: any) => void;
}

const decorate = withStyles(
  (theme: any) => ({
    root: {
      marginBottom: theme.spacing * 2,
      display: "flex" as "flex"
    },
    spacer: {
      flex: 1
    }
  }),
  { displayName: "FormSave" }
);
export const FormSave = decorate<Props>(
  ({ classes, disabled, variant, onConfirm, ...props }) => (
    <div {...props}>
      <hr />
      <div className={classes.root}>
        <div className={classes.spacer} />
        <Button
          bsStyle={
            variant === "success"
              ? "success"
              : variant === "error"
                ? "danger"
                : variant === "loading"
                  ? "primary"
                  : "primary"
          }
          disabled={variant === "loading" || disabled}
          onClick={onConfirm}
        >
          {variant === "success"
            ? i18n.t("Saved")
            : variant === "error"
              ? i18n.t("Error")
              : variant === "loading"
                ? i18n.t("Loading")
                : i18n.t("Save")}
        </Button>
      </div>
    </div>
  )
);
export default FormSave;
