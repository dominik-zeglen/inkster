import * as React from "react";
import { Panel } from "react-bootstrap";

import Input from '../../components/Input'
import i18n from '../../i18n'

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
  onChange
}) => (
  <Panel>
    <Panel.Heading>
      <Panel.Title>{i18n.t("General information")}</Panel.Title>
    </Panel.Heading>
    <Panel.Body>
      <Input
        disabled={disabled}
        name="email"
        value={data.email}
        onChange={onChange}
      />
    </Panel.Body>
  </Panel>
);
export default UserProperties;
