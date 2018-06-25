import * as React from "react";
import { Plus } from "react-feather";

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
  onNextPage?: () => void;
  onPreviousPage?: () => void;
}

export const ContainerListPage: React.StatelessComponent<Props> = ({
  containers,
  container,
  disabled,
  hasNextPage,
  hasPreviousPage,
  variant,
  onBack,
  onContainerAdd,
  onNextPage,
  onPreviousPage
}) => (
  <Container width="md">
    <PageHeader
      title={
        variant === "root"
          ? i18n.t("Containers")
          : container
            ? container.name
            : undefined
      }
      onBack={onBack}
    >
      <IconButton disabled={disabled} icon={Plus} onClick={onContainerAdd} />
    </PageHeader>
    <ContainerList
      containers={containers}
      hasNextPage={hasNextPage}
      hasPreviousPage={hasPreviousPage}
      onPreviousPage={onPreviousPage}
      onNextPage={onNextPage}
    />
  </Container>
);
export default ContainerListPage;
