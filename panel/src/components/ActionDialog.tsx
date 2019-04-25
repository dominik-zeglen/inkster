import * as React from "react";
import withStyles from "react-jss";
import Button from "aurora-ui-kit/dist/components/Button";
import Dialog from "aurora-ui-kit/dist/components/Dialog";
import DialogActions from "aurora-ui-kit/dist/components/DialogActions";
import DialogContent from "aurora-ui-kit/dist/components/DialogContent";
import DialogHeader from "aurora-ui-kit/dist/components/DialogHeader";
import IconButton from "aurora-ui-kit/dist/components/IconButton";
import { X } from "react-feather";

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
          maxWidth: theme.breakpoints.width(current),
        },
      };
      return prev;
    }, {}),
  { displayName: "ActionDialog" },
);
export const ActionDialog = decorate<Props>(
  ({ children, show, size, title, onClose, onConfirm }) => {
    const handleSubmit = (event: any) => {
      onConfirm(event);
      onClose();
    };
    return (
      <Dialog isOpen={show} onClose={onClose} size={size}>
        <DialogHeader title={title}>
          <IconButton onClick={onClose}>
            <X />
          </IconButton>
        </DialogHeader>
        <DialogContent>{children}</DialogContent>
        <DialogActions>
          <Button variant="outlined" onClick={onClose}>
            {i18n.t("Close")}
          </Button>
          <Button variant="default" onClick={handleSubmit}>
            {i18n.t("Confirm")}
          </Button>
        </DialogActions>
      </Dialog>
    );
  },
);
export default ActionDialog;
