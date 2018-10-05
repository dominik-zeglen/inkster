import * as React from "react";

export interface LinkProps extends React.HTMLProps<HTMLAnchorElement> {
  onClick: (event?: React.MouseEvent<any>) => void;
}

export const Link: React.StatelessComponent<LinkProps> = ({
  onClick,
  ...linkProps
}) => {
  const handleClick = (event: React.MouseEvent<any>) => {
    event.preventDefault();
    event.stopPropagation();
    onClick(event);
  };
  return <a {...linkProps} onClick={handleClick} />;
};
