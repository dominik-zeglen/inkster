import * as React from "react";
import { Panel } from "react-bootstrap";
import { X } from "react-feather";

import IconButton from "../../components/IconButton";
import Input from "../../components/Input";
import RichTextEditor from "../../components/RichTextEditor";
import i18n from "../../i18n";

interface Props {
  data: {
    id: string;
    name: string;
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
      <IconButton icon={X} onClick={onDelete} />
    </Panel.Heading>
    <Panel.Body>
      <Input
        label={i18n.t("Name")}
        name="name"
        value={data.name}
        onChange={onChange}
      />
      {data.type === "longText" ? (
        <RichTextEditor
          label={i18n.t("Field value")}
          initialValue={data.value}
          name="value"
          onChange={onChange}
        />
      ) : (
        <Input
          label={i18n.t("Value")}
          name="value"
          value={data.value}
          onChange={onChange}
        />
      )}
    </Panel.Body>
  </Panel>
);

export default PageFieldProperties;
