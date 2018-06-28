import * as React from "react";
import { Panel } from "react-bootstrap";

import Input from "../../components/Input";
import i18n from "../../i18n";

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
  onChange
}) => (
  <Panel>
    <Panel.Body>
      <Input
        label={i18n.t("Name")}
        name="name"
        placeholder={i18n.t("Name")}
        value={data.name}
        onChange={onChange}
      />
    </Panel.Body>
  </Panel>
);
export default DirectoryProperties;
