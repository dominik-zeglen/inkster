import { storiesOf } from "@storybook/react";

import { makeProps, makeStories, listViewProps } from "../fixtures";
import { users } from "./fixtures";
import Decorator from "../Decorator";
import UserListPage from "../../src/users/components/UserListPage";

const stories = storiesOf("Views / Users / User list", module).addDecorator(
  Decorator
);

const props = makeProps(listViewProps, {
  default: { users },
  noData: { users: [] },
  loading: {}
});
makeStories(stories, props, UserListPage);
