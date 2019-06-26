import * as React from "react";
import Button from "aurora-ui-kit/dist/components/Button";

import { StandardProps } from "./";
import { TransactionState } from "../";
import i18n from "../i18n";
import createUseStyles from "aurora-ui-kit/dist/utils/jss";
import { ITheme } from "aurora-ui-kit/dist/theme";

interface Props extends StandardProps {
  disabled: boolean;
  variant: TransactionState;
  onConfirm: (event: any) => void;
}

const useStyles = createUseStyles((theme: ITheme) => ({
  hr: {
    height: 1,
    background: theme.mixins.fade(theme.colors.gray.main, theme.alpha.default),
    border: "none",
  },
  root: {
    marginBottom: theme.spacing * 2,
    display: "flex" as "flex",
  },
  spacer: {
    flex: 1,
  },
}));
export const FormSave: React.FC<Props> = ({
  disabled,
  variant,
  onConfirm,
  ...props
}) => {
  const classes = useStyles();

  return (
    <div {...props}>
      <hr className={classes.hr} />
      <div className={classes.root}>
        <div className={classes.spacer} />
        <Button
          color={
            variant === "success"
              ? "success"
              : variant === "error"
              ? "error"
              : variant === "loading"
              ? "secondary"
              : "secondary"
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
  );
};
export default FormSave;
