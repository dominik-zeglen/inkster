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

export const PageProperties: React.StatelessComponent<Props> = ({
  data,
  disabled,
  onChange
}) => (
  <Panel>
    <Panel.Body>
      <Input
        name="name"
        disabled={disabled}
        label={i18n.t("Name")}
        value={data.name}
        onChange={onChange}
      />
    </Panel.Body>
  </Panel>
);
export default PageProperties;
