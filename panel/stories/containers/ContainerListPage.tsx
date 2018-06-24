import * as React from "react";

import { storiesOf } from "@storybook/react";
import { action } from "@storybook/addon-actions";

import Decorator from "../Decorator";
import ContainerListPage from "../../src/containers/components/ContainerListPage";
import { containers, container } from "./fixtures";

storiesOf("Containers / Container list / Root", module)
  .addDecorator(Decorator)
  .add("default", () => (
    <ContainerListPage
      containers={containers}
      hasPreviousPage={false}
      hasNextPage={true}
      variant="root"
      onNextPage={action("next")}
      onPreviousPage={action("previous")}
    />
  ))
  .add("loading", () => (
    <ContainerListPage
      hasPreviousPage={false}
      hasNextPage={true}
      variant="root"
      onNextPage={action("next")}
      onPreviousPage={action("previous")}
    />
  ));
storiesOf("Containers / Container list / Child container", module)
  .addDecorator(Decorator)
  .add("default", () => (
    <ContainerListPage
      containers={containers}
      container={container}
      variant="child"
      hasPreviousPage={false}
      hasNextPage={true}
      onBack={action("back")}
      onNextPage={action("next")}
      onPreviousPage={action("previous")}
    />
  ))
  .add("loading", () => (
    <ContainerListPage
      variant="child"
      hasPreviousPage={false}
      hasNextPage={true}
      onBack={action("back")}
      onNextPage={action("next")}
      onPreviousPage={action("previous")}
    />
  ));
