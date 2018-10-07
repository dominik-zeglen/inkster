import { storiesOf } from "@storybook/react";
import * as React from "react";

import Decorator from "../Decorator";
import PasswordResetPage, { Props } from "../../src/auth/components/PasswordResetPage";

const props: Props = {
  disabled: false,
  onSubmit: () => undefined
};

storiesOf("Views / Authentication / Password reset", module)
  .addDecorator(Decorator)
  .add("default", () => <PasswordResetPage {...props} />)
