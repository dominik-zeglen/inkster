import { storiesOf } from "@storybook/react";
import * as React from "react";

import Decorator from "../Decorator";
import {
  Notification,
  NotificationProps,
  NotificationType
} from "../../src/components/Notificator/Notification";

const props: NotificationProps = {
  onClose: () => console.log("notification should close"),
  text: "Lorem ipsum dolor sit amet",
  type: NotificationType.DEFAULT
};

storiesOf("Components / Notifications", module)
  .addDecorator(Decorator)
  .add("default", () => <Notification {...props} />)
  .add("warning", () => (
    <Notification {...props} type={NotificationType.WARNING} />
  ))
  .add("error", () => (
    <Notification {...props} type={NotificationType.ERROR} />
  ));
