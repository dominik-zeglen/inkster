import * as React from "react";
import {
  ControlLabel,
  FormControl,
  FormGroup,
  HelpBlock
} from "react-bootstrap";

import { StandardProps } from "./";

interface Props extends StandardProps {
  disabled?: boolean;
  error?: boolean;
  helperText?: string;
  id?: string;
  label?: string;
  name: string;
  placeholder?: string;
  type?: string;
  value: string;
  onChange: (event: React.ChangeEvent<any>) => void;
}

export const Input: React.StatelessComponent<Props> = ({
  error,
  helperText,
  id,
  label,
  name,
  onChange,
  placeholder,
  type,
  value
}) => (
  <FormGroup controlId={id} validationState={error ? "error" : null}>
    {label && <ControlLabel>{label}</ControlLabel>}
    <FormControl
      name={name}
      type={type || "text"}
      value={value}
      onChange={onChange}
      placeholder={placeholder}
    />
    {helperText && <HelpBlock>{helperText}</HelpBlock>}
  </FormGroup>
);
export default Input;
