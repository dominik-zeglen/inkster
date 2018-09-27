import * as React from "react";
import { Checkbox as BsCheckbox } from "react-bootstrap";

interface Props {
  label: string;
  name: string;
  value: boolean;
  onChange: (event: React.ChangeEvent<any>) => void;
}

export const Checkbox: React.StatelessComponent<Props> = ({
  label,
  value,
  onChange,
  ...props
}) => {
  const handleCheckboxChange = (event: React.ChangeEvent<any>) => {
    onChange({
      target: {
        name: event.target.name,
        value: event.target.checked
      }
    } as any);
  };
  return (
    <BsCheckbox checked={value} onChange={handleCheckboxChange} {...props}>
      {label}
    </BsCheckbox>
  );
};
export default Checkbox;
