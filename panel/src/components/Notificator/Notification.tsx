import * as React from "react";
import { Alert } from "react-bootstrap";
import { AlertTriangle, Info } from "react-feather";
import withStyles from "react-jss";

export enum NotificationType {
  DEFAULT,
  WARNING,
  ERROR,
}

const CLOSE_AFTER = 5000;
const ICON_SIZE = 20;

interface NotificationComponentProps {
  text: string;
  type?: NotificationType;
  onDismiss: () => void;
  onPointerEnter: () => void;
  onPointerLeave: () => void;
}
const decorate = withStyles(theme => ({
  iconContainer: {
    alignItems: "center" as "center",
    display: "flex" as "flex",
  },
  root: {
    "& svg": {
      height: ICON_SIZE,
      width: ICON_SIZE,
    },
    display: "grid" as "grid",
    gridColumnGap: theme.spacing * 2 + "px",
    gridTemplateColumns: `${ICON_SIZE}px 1fr`,
    zIndex: 1,
  },
}));

const NotificationComponent = decorate<NotificationComponentProps>(
  ({ classes, onDismiss, onPointerEnter, onPointerLeave, text, type }) => (
    <Alert
      bsStyle={
        type === NotificationType.ERROR
          ? "danger"
          : type === NotificationType.WARNING
            ? "warning"
            : "info"
      }
      onDismiss={onDismiss}
      onPointerEnter={onPointerEnter}
      onPointerLeave={onPointerLeave}
    >
      <div className={classes.root}>
        <div className={classes.iconContainer}>
          {type === NotificationType.ERROR ? (
            <AlertTriangle />
          ) : type === NotificationType.WARNING ? (
            <AlertTriangle />
          ) : (
            <Info />
          )}
        </div>
        <div>{text}</div>
      </div>
    </Alert>
  ),
);

export interface NotificationProps {
  closeAfter?: number;
  text: string;
  type?: NotificationType;
  onClose: () => void;
}
interface NotificationState {
  timer: any | null;
}
export class Notification extends React.Component<
  NotificationProps,
  NotificationState
> {
  state: NotificationState = {
    timer: setTimeout(this.props.onClose, this.props.closeAfter || CLOSE_AFTER),
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
        this.props.closeAfter || CLOSE_AFTER,
      ),
    });

  render() {
    const { text, type, onClose } = this.props;

    return (
      <NotificationComponent
        onDismiss={onClose}
        onPointerEnter={this.handlePointerEnter}
        onPointerLeave={this.handlePointerLeave}
        text={text}
        type={type}
      />
    );
  }
}
export default Notification;
