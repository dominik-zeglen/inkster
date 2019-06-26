import * as React from "react";
import { AlertTriangle, Info } from "react-feather";
import { ITheme } from "aurora-ui-kit/dist/theme";
import createUseStyles from "aurora-ui-kit/dist/utils/jss";
import NotificationInnerComponent from "aurora-ui-kit/dist/components/Notification";

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
const useStyles = createUseStyles((theme: ITheme) => ({
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

const NotificationComponent: React.FC<NotificationComponentProps> = ({
  onDismiss,
  onPointerEnter,
  onPointerLeave,
  text,
  type,
}) => {
  const classes = useStyles();
  return (
    <NotificationInnerComponent
      color={
        type === NotificationType.ERROR
          ? "error"
          : type === NotificationType.WARNING
          ? "warning"
          : "primary"
      }
      componentProps={{
        onMouseEnter: onPointerEnter,
        onMouseLeave: onPointerLeave,
      }}
      onClick={onDismiss}
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
    </NotificationInnerComponent>
  );
};

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
