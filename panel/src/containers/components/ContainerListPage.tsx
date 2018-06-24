import * as React from "react";

import Container from "../../components/Container";
import PageHeader from "../../components/PageHeader";
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
  hasNextPage?: boolean;
  hasPreviousPage?: boolean;
  variant: "root" | "child";
  onBack?: () => void;
  onNextPage?: () => void;
  onPreviousPage?: () => void;
}

export const ContainerListPage: React.StatelessComponent<Props> = ({
  containers,
  container,
  hasNextPage,
  hasPreviousPage,
  variant,
  onBack,
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
    />
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
