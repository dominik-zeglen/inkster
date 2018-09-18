import * as React from "react";
import { Modal as ModalComponent, ModalProps } from "react-bootstrap";
import withStyles from "react-jss";

interface Props extends ModalProps {
  width: "xs" | "sm" | "md" | "lg";
}

const decorate = withStyles(
  theme =>
    ["xs", "sm", "md", "lg"].reduce((prev, current) => {
      prev[current] = {
        [theme.breakpoints.up(current)]: {
          marginLeft: "auto",
          marginRight: "auto",
          maxWidth: theme.breakpoints.width(current)
        }
      };
      return prev;
    }, {}),
  { displayName: "Modal" }
);

export const Modal = decorate<Props>(({ classes, width, ...props }) => (
  <ModalComponent className={classes[width]} {...props} />
));
export default Modal;
