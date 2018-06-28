import { storiesOf } from "@storybook/react";

import { makeProps, makeStories, listViewProps } from "../fixtures";
import { directories } from "./fixtures";
import Decorator from "../Decorator";
import DirectoryRootPage from "../../src/directories/components/DirectoryRootPage";

const stories = storiesOf(
  "Views / Directories / Directory root",
  module
).addDecorator(Decorator);

const props = makeProps(listViewProps, {
  default: { directories },
  noData: { directories: [] },
  loadingData: {}
});
makeStories(stories, props, DirectoryRootPage);
