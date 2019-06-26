import * as React from "react";
import createUseStyles from "aurora-ui-kit/dist/utils/jss";
import Button from "aurora-ui-kit/dist/components/Button";
import Input from "aurora-ui-kit/dist/components/TextInput";

import PageLayout from "./PageLayout";
import Form from "../../components/Form";
import i18n from "../../i18n";

export interface FormData {
  email: string;
}

export interface Props {
  disabled: boolean;
  onSubmit: (data: FormData) => void;
}

const initialForm: FormData = {
  email: "",
};
const useStyles = createUseStyles({
  buttonContainer: {
    display: "flex" as "flex",
    justifyContent: "flex-end" as "flex-end",
  },
});
export const PasswordResetSendEmailPage: React.FC<Props> = ({
  disabled,
  onSubmit,
}) => {
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
            label={i18n.t("E-mail")}
            helpText={i18n.t(
              "Enter your e-mail address, so we can send you message with further instructions",
              {
                context: "caption",
              },
            )}
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
          <div className={classes.buttonContainer}>
            <Button
              color="primary"
              disabled={disabled || !hasChanged}
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
export default PasswordResetSendEmailPage;
