import * as React from "react";
import { storiesOf } from "@storybook/react";

import Decorator from "../Decorator";
import { viewer } from "./fixtures";
import HomePage, { Props } from "../../src/home/components/HomePage";

const props: Props = {
  disabled: false,
  onPageClick: () => undefined,
  user: viewer,
};

storiesOf("Views / Home page", module)
  .addDecorator(Decorator)
  .add("default", () => <HomePage {...props} />)
  .add("disabled", () => <HomePage {...props} disabled={true} />);
