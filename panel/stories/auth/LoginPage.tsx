import { storiesOf } from "@storybook/react";
import * as React from "react";

import Decorator from "../Decorator";
import LoginPage, { Props } from "../../src/auth/components/LoginPage";

const props: Props = {
  disabled: false,
  error: false,
  passwordRecoveryHref: '#',
  onSubmit: () => undefined
};

storiesOf("Views / Authentication / Login page", module)
  .addDecorator(Decorator)
  .add("default", () => <LoginPage {...props} />)
  .add("error", () => <LoginPage {...props} error={true} />)
