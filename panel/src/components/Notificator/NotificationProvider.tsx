import * as React from "react";
import withStyles from "react-jss";

import Notification, {
  NotificationProps,
  NotificationType
} from "./Notification";
import { NotificatorContext } from ".";

type Omit<T, K> = Pick<T, Exclude<keyof T, K>>;
interface NotificationProviderState {
  notifications: Array<
    Omit<NotificationProps, "onClose"> & {
      key: string;
    }
  >;
}

const decorate = withStyles(theme => ({
  notificationContainer: {
    bottom: theme.spacing * 2,
    marginLeft: theme.spacing * 2,
    maxWidth: 340,
    position: "fixed" as "fixed",
    right: theme.spacing * 2,
    width: '100%'
  }
}));
const NotificationContainer = decorate(({ classes, ...props }) => (
  <div className={classes.notificationContainer} {...props} />
));

export class NotificationProvider extends React.Component<
  {},
  NotificationProviderState
> {
  state: NotificationProviderState = {
    notifications: []
  };

  handleNotificationClose = (key: string) =>
    this.setState({
      notifications: this.state.notifications.filter(
        notification => notification.key !== key
      )
    });

  handleNotificationPush = (notification: {
    text: string;
    type?: NotificationType;
  }) =>
    this.setState({
      notifications: this.state.notifications.concat({
        ...notification,
        key: new Date().getTime().toString()
      })
    });

  render() {
    const { notifications } = this.state;

    return (
      <>
        <NotificationContainer>
          {notifications.map(notification => (
            <Notification
              {...notification}
              onClose={() => this.handleNotificationClose(notification.key)}
            />
          ))}
        </NotificationContainer>
        <NotificatorContext.Provider value={this.handleNotificationPush}>
          {this.props.children}
        </NotificatorContext.Provider>
      </>
    );
  }
}
export default NotificationProvider;
