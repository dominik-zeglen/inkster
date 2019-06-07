import { ISelectOption } from "aurora-ui-kit/dist/components/Select";

import i18n from "../i18n";

export const fieldTypes: () => ISelectOption[] = () => [
  {
    label: i18n.t("Short text"),
    value: "text",
  },
  {
    label: i18n.t("Long text"),
    value: "longText",
  },
  {
    label: i18n.t("Image"),
    value: "image",
  },
  {
    label: i18n.t("File"),
    value: "file",
  },
];
