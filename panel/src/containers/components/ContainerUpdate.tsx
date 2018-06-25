import * as React from "react";
import { Plus, Trash } from "react-feather";

import Container from "../../components/Container";
import PageHeader from "../../components/PageHeader";
import IconButton from "../../components/IconButton";
import ContainerList from "./ContainerList";
import i18n from "../../i18n";


interface Props {
  containers?: Array<{
    id?: string,
    name?: string
  }>;
  container?: {
    id?: string,
    name?: string
  };
  disabled?: boolean;
  hasNextPage?: boolean;
  hasPreviousPage?: boolean;
  variant: "root" | "child";
  onBack?: () => void;
  onContainerAdd?: () => void;
  onSubcontainersNextPage?: () => void;
  onSubcontainersPreviousPage?: () => void;
}
