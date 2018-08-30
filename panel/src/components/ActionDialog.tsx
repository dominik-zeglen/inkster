import * as React from "react";
import { Button, Modal } from "react-bootstrap";
import withStyles from "react-jss";

import i18n from "../i18n";

interface Props {
  show: boolean;
  size: "xs" | "sm" | "md" | "lg";
  title: string;
  onClose: () => void;
  onConfirm: (args: any) => void;
}

const decorate = withStyles(
  (theme: any) =>
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
  { displayName: "ActionDialog" }
);
export const ActionDialog = decorate<Props>(
  ({ classes, children, show, size, title, onClose, onConfirm }) => {
    const handleSubmit = (event: any) => {
      onConfirm(event);
      onClose();
    };
    return (
      <Modal show={show} onHide={onClose} className={classes[size]}>
        <Modal.Header>
          <Modal.Title>{title}</Modal.Title>
        </Modal.Header>
        <Modal.Body>{children}</Modal.Body>
        <Modal.Footer>
          <Button onClick={onClose}>{i18n.t("Close")}</Button>
          <Button bsStyle="primary" onClick={handleSubmit} type="submit">
            {i18n.t("Confirm")}
          </Button>
        </Modal.Footer>
      </Modal>
    );
  }
);
export default ActionDialog;
