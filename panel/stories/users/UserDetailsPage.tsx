import { storiesOf } from "@storybook/react";

import { formViewProps, makeProps, makeStories } from "../fixtures";
import { user } from "./fixtures";
import Decorator from "../Decorator";
import UserDetailsPage from "../../src/users/components/UserDetailsPage";

const stories = storiesOf("Views / Users / User details", module).addDecorator(
  Decorator
);

const props = makeProps(formViewProps, {
  default: { user },
  loading: { user: undefined }
});
makeStories(stories, props, UserDetailsPage);
