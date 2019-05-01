import * as React from "react";
import Card from "aurora-ui-kit/dist/components/Card";
import CardContent from "aurora-ui-kit/dist/components/CardContent";
import CardHeader from "aurora-ui-kit/dist/components/CardHeader";
import CardTitle from "aurora-ui-kit/dist/components/CardTitle";
import Input from "aurora-ui-kit/dist/components/TextInput";

import i18n from "../../i18n";

interface Props {
  data: {
    email: string;
  };
  disabled: boolean;
  onChange: (event: React.ChangeEvent<any>) => void;
}

export const UserProperties: React.StatelessComponent<Props> = ({
  data,
  disabled,
  onChange,
}) => (
  <Card>
    <CardHeader>
      <CardTitle>{i18n.t("General Informations")}</CardTitle>
    </CardHeader>
    <CardContent>
      <Input
        disabled={disabled}
        label={i18n.t("User e-mail")}
        value={data.email}
        onChange={value =>
          onChange({
            target: {
              name: "email",
              value,
            },
          } as any)
        }
      />
    </CardContent>
  </Card>
);
export default UserProperties;
