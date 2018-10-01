import { storiesOf } from "@storybook/react";
import * as React from "react";

import Decorator from "../Decorator";
import PasswordResetSendEmailPage, {
  Props
} from "../../src/auth/components/PasswordResetSendEmailPage";

const props: Props = {
  disabled: false,
  onSubmit: () => undefined
};

storiesOf("Views / Authentication / Password reset send email page", module)
  .addDecorator(Decorator)
  .add("default", () => <PasswordResetSendEmailPage {...props} />);
