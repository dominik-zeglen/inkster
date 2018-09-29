import * as React from "react";
import { Alert } from "react-bootstrap";

export enum NotificationType {
  DEFAULT,
  WARNING,
  ERROR
}

export interface NotificationProps {
  closeAfter?: number;
  text: string;
  type?: NotificationType;
  onClose: () => void;
}
interface NotificationState {
  timer: any | null;
}

const CLOSE_AFTER = 50000;

export class Notification extends React.Component<
  NotificationProps,
  NotificationState
> {
  state: NotificationState = {
    timer: setTimeout(this.props.onClose, this.props.closeAfter || CLOSE_AFTER)
  };

  componentWillUnmount() {
    if (this.state.timer) {
      clearTimeout(this.state.timer);
    }
  }

  handlePointerEnter = () => {
    const { timer } = this.state;
    if (timer) {
      clearTimeout(timer);
      this.setState({ timer: null });
    }
  };

  handlePointerLeave = () =>
    this.setState({
      timer: setTimeout(
        this.props.onClose,
        this.props.closeAfter || CLOSE_AFTER
      )
    });

  render() {
    const { text, type, onClose } = this.props;

    return (
      <Alert
        bsStyle={
          type === NotificationType.ERROR
            ? "danger"
            : type === NotificationType.WARNING
              ? "warning"
              : "info"
        }
        onDismiss={onClose}
        onPointerEnter={this.handlePointerEnter}
        onPointerLeave={this.handlePointerLeave}
      >
        {text}
      </Alert>
    );
  }
}
export default Notification;
