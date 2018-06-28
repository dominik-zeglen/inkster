import { storiesOf } from "@storybook/react";

import {
  makeProps,
  makeStories,
  listViewProps,
  formViewProps
} from "../fixtures";
import { directory } from "./fixtures";
import Decorator from "../Decorator";
import DirectoryDetailsPage from "../../src/directories/components/DirectoryDetailsPage";

const stories = storiesOf(
  "Views / Directories / Directory details",
  module
).addDecorator(Decorator);

const props = makeProps(makeProps(listViewProps, formViewProps), {
  default: { directory },
  loading: { directory: undefined },
});
makeStories(stories, props, DirectoryDetailsPage);
