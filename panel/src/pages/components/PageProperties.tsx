import * as React from "react";
import { Panel } from "react-bootstrap";

import Input from "../../components/Input";
import i18n from "../../i18n";

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
  onChange
}) => (
  <Panel>
    <Panel.Heading>
      <Panel.Title>{i18n.t("General information")}</Panel.Title>
    </Panel.Heading>
    <Panel.Body>
      <Input
        name="name"
        disabled={disabled}
        label={i18n.t("Name")}
        value={data.name}
        onChange={onChange}
      />
      <Input
        name="slug"
        disabled={disabled}
        label={i18n.t("Slug")}
        value={data.slug}
        onChange={onChange}
      />
    </Panel.Body>
  </Panel>
);
export default PageProperties;
