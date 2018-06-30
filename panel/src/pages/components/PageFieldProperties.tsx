import * as React from "react";
import { Panel } from "react-bootstrap";
import { Trash } from "react-feather";

import IconButton from "../../components/IconButton";
import Input from "../../components/Input";
import i18n from "../../i18n";

interface Props {
  data: {
    id: string;
    name: string;
    slug: string;
    type: string;
    value: string;
  };
  name: string;
  onChange: (event: React.ChangeEvent<any>) => void;
  onDelete: () => void;
}

export const PageFieldProperties: React.StatelessComponent<Props> = ({
  data,
  name,
  onChange,
  onDelete,
  ...props
}) => (
  <Panel {...props}>
    <Panel.Heading>
      <Panel.Title>{i18n.t("Field properties")}</Panel.Title>
      <IconButton icon={Trash} onClick={onDelete} />
    </Panel.Heading>
    <Panel.Body>
      <Input
        label={i18n.t("Field name")}
        name="name"
        value={data.name}
        onChange={onChange}
      />
      <Input
        label={i18n.t("Field slug")}
        name="slug"
        value={data.slug}
        onChange={onChange}
      />
      <Input
        label={i18n.t("Field value")}
        name="value"
        value={data.value}
        type={data.type === "longText" ? "textarea" : "text"}
        onChange={onChange}
      />
    </Panel.Body>
  </Panel>
);

export default PageFieldProperties;
