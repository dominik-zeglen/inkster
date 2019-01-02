import * as React from "react";
import { Panel } from "react-bootstrap";
import { FileText } from "react-feather";

import i18n from "../../i18n";
import { renderCollection, maybe } from "../../utils";
import ListElement from "../../components/ListElement";
import { Viewer_viewer_pages_edges_node } from "../queries/types/Viewer";

interface Props {
  disabled: boolean;
  pages: Viewer_viewer_pages_edges_node[];
  onPageClick: (id: string) => void;
}

export const HomePageNewestPages: React.StatelessComponent<Props> = ({
  disabled,
  pages,
  onPageClick,
}) => (
  <Panel>
    <Panel.Heading>
      <Panel.Title>{i18n.t("Your newest pages")}</Panel.Title>
    </Panel.Heading>
    <Panel.Body>
      {renderCollection(pages, page => (
        <ListElement
          icon={FileText}
          disabled={disabled}
          title={maybe(() => page.name)}
          onClick={page ? () => onPageClick(page.id) : undefined}
        />
      ))}
    </Panel.Body>
  </Panel>
);
export default HomePageNewestPages;
