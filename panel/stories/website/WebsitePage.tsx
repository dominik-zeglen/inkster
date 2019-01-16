import { storiesOf } from "@storybook/react";

import { formViewProps, makeProps, makeStories } from "../fixtures";
import { website } from "./fixtures";
import Decorator from "../Decorator";
import WebsitePage from "../../src/website/components/WebsitePage";

const stories = storiesOf("Views / Website settings", module).addDecorator(
  Decorator,
);

const props = makeProps(formViewProps, {
  default: { website },
  loading: { website: undefined },
});
makeStories(stories, props, WebsitePage);
