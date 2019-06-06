import * as React from "react";
import Card from "aurora-ui-kit/dist/components/Card";
import CardHeader from "aurora-ui-kit/dist/components/CardHeader";
import CardTitle from "aurora-ui-kit/dist/components/CardTitle";
import CardContent from "aurora-ui-kit/dist/components/CardContent";
import Input from "aurora-ui-kit/dist/components/TextInput";

import i18n from "../../i18n";
import Spacer from "../../components/Spacer";

interface Props {
  data: {
    name: string;
    slug: string;
  };
  disabled: boolean;
  onChange: (event: React.ChangeEvent<any>) => void;
}

export const PageProperties: React.StatelessComponent<Props> = ({
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
        label={i18n.t("Name")}
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
      <Spacer />
      <Input
        disabled={disabled}
        label={i18n.t("Slug")}
        value={data.slug}
        onChange={value =>
          onChange({
            target: {
              name: "slug",
              value,
            },
          } as any)
        }
      />
    </CardContent>
  </Card>
);
export default PageProperties;
