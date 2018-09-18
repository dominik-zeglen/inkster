import * as React from "react";
import { Button, Modal as BsModal } from "react-bootstrap";

import i18n from "../i18n";
import Form, { FormChildren } from "./Form";
import Modal from "./Modal";

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
      onConfirm
    } = this.props;
    const handleSubmit = (data: T) => {
      onConfirm(data);
      onClose();
    };
    return (
      <Modal show={show} onHide={onClose} width={width}>
        <Form initial={initial} onSubmit={handleSubmit}>
          {formData => (
            <>
              <BsModal.Header>
                <BsModal.Title>{title}</BsModal.Title>
              </BsModal.Header>
              <BsModal.Body>
                {typeof children === "function" ? children(formData) : children}
              </BsModal.Body>
              <BsModal.Footer>
                <Button onClick={onClose}>{i18n.t("Close")}</Button>
                <Button
                  bsStyle="primary"
                  onClick={formData.submit}
                  type="submit"
                >
                  {i18n.t("Confirm")}
                </Button>
              </BsModal.Footer>
            </>
          )}
        </Form>
      </Modal>
    );
  }
}

export default FormDialog;
