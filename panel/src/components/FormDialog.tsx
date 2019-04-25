import * as React from "react";
import Button from "aurora-ui-kit/dist/components/Button";
import Dialog from "aurora-ui-kit/dist/components/Dialog";
import DialogActions from "aurora-ui-kit/dist/components/DialogActions";
import DialogContent from "aurora-ui-kit/dist/components/DialogContent";
import DialogHeader from "aurora-ui-kit/dist/components/DialogHeader";
import IconButton from "aurora-ui-kit/dist/components/IconButton";
import { X } from "react-feather";

import i18n from "../i18n";
import Form, { FormChildren } from "./Form";

interface Props<T extends {}> {
  children: FormChildren<T>;
  show: boolean;
  width: "xs" | "sm" | "md" | "lg";
  title: string;
  initial: T;
  onClose: () => void;
  onConfirm: (data?: T) => void;
}

export class FormDialog<T extends {} = {}> extends React.Component<
  Props<T>,
  {}
> {
  render() {
    const {
      children,
      initial,
      show,
      width,
      title,
      onClose,
      onConfirm,
    } = this.props;
    return (
      <Dialog isOpen={show} onClose={onClose} size={width}>
        <Form initial={initial} onSubmit={onConfirm}>
          {formData => (
            <>
              <DialogHeader title={title}>
                <IconButton onClick={onClose}>
                  <X />
                </IconButton>
              </DialogHeader>
              <DialogContent>{children(formData)}</DialogContent>
              <DialogActions>
                <Button
                  componentProps={
                    {
                      type: "button",
                    } as any
                  }
                  variant="outlined"
                  onClick={onClose}
                >
                  {i18n.t("Close")}
                </Button>
                <Button
                  componentProps={
                    {
                      type: "submit",
                    } as any
                  }
                  variant="default"
                  onClick={formData.submit}
                >
                  {i18n.t("Confirm")}
                </Button>
              </DialogActions>
            </>
          )}
        </Form>
      </Dialog>
    );
  }
}

export default FormDialog;
