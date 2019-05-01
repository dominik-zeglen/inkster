import * as React from "react";
import Card from "aurora-ui-kit/dist/components/Card";
import CardContent from "aurora-ui-kit/dist/components/CardContent";
import Input from "aurora-ui-kit/dist/components/TextInput";

import i18n from "../../i18n";
import CardHeader from "aurora-ui-kit/dist/components/CardHeader";
import CardTitle from "aurora-ui-kit/dist/components/CardTitle";

interface Props {
  data: {
    name: string;
  };
  disabled: boolean;
  onChange: (event: React.ChangeEvent<any>) => void;
}

export const DirectoryProperties: React.StatelessComponent<Props> = ({
  disabled,
  data,
  onChange,
}) => (
  <Card>
    <CardHeader>
      <CardTitle>{i18n.t("General Informations")}</CardTitle>
    </CardHeader>
    <CardContent>
      <Input
        disabled={disabled}
        label={i18n.t("Name")}
        placeholder={i18n.t("Name")}
        value={data.name}
        onChange={value =>
          onChange({
            target: {
              name: "name",
              value,
            },
          } as any)
        }
      />
    </CardContent>
  </Card>
);
export default DirectoryProperties;
