import * as React from "react";
import createUseStyles from "aurora-ui-kit/dist/utils/jss";
import Button from "aurora-ui-kit/dist/components/Button";
import Input from "aurora-ui-kit/dist/components/TextInput";

import PageLayout from "./PageLayout";
import Form from "../../components/Form";
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
  passwordConfirm: "",
};
const useStyles = createUseStyles({
  buttonContainer: {
    display: "flex" as "flex",
    justifyContent: "flex-end" as "flex-end",
  },
});
export const PasswordResetPage: React.FC<Props> = ({ disabled, onSubmit }) => {
  const classes = useStyles();

  return (
    <Form initial={initialForm} onSubmit={onSubmit}>
      {({ change, data, hasChanged }) => (
        <PageLayout
          header={i18n.t("Reset password", {
            context: "header",
          })}
        >
          <Input
            label={i18n.t("New password", {
              context: "label",
            })}
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
          <Input
            error={data.password !== data.passwordConfirm}
            label={i18n.t("Confirm password", {
              context: "label",
            })}
            helpText={
              data.password !== data.passwordConfirm
                ? i18n.t("Passwords do not match", {
                    context: "caption",
                  })
                : undefined
            }
            InputProps={{
              componentProps: {
                type: "password",
              },
            }}
            value={data.passwordConfirm}
            onChange={value =>
              change({
                target: {
                  name: "passwordConfirm",
                  value,
                },
              } as any)
            }
          />
          <div className={classes.buttonContainer}>
            <Button
              color="primary"
              disabled={
                disabled ||
                !hasChanged ||
                data.password !== data.passwordConfirm
              }
              type="submit"
            >
              {i18n.t("Submit", { context: "button" })}
            </Button>
          </div>
        </PageLayout>
      )}
    </Form>
  );
};
export default PasswordResetPage;
