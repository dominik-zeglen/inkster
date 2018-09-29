import * as React from "react";

import { NotificationType } from "./Notification";

type NotificatorContextType = (
  notification: { type?: NotificationType; text: string }
) => void;

export const NotificatorContext = React.createContext<NotificatorContextType>(
  () => undefined
);

export * from "./Notification";
export * from "./NotificationProvider";
export default NotificatorContext.Consumer;
