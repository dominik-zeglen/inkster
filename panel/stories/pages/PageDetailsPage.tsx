import { storiesOf } from "@storybook/react";

import { makeProps, makeStories, formViewProps } from "../fixtures";
import { page } from "./fixtures";
import Decorator from "../Decorator";
import PageDetailsPage from "../../src/pages/components/PageDetailsPage";

const stories = storiesOf("Views / Pages / Page details", module).addDecorator(
  Decorator
);

const props = makeProps(formViewProps, {
  default: { page },
  loading: { page: undefined }
});
makeStories(stories, props, PageDetailsPage);
