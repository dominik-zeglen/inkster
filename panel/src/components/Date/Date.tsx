import * as React from "react";
import * as moment from "moment";
import { OverlayTrigger, Tooltip } from "react-bootstrap";

import { Consumer } from "./DateContext";

interface Props {
  date: string;
}

export const Date: React.StatelessComponent<Props> = ({ date }) => (
  <Consumer>
    {dateNow => (
      <OverlayTrigger
        placement="bottom"
        overlay={
          <Tooltip id="id">
            {moment(date)
              .toDate()
              .toLocaleString()}
          </Tooltip>
        }
      >
        <time>{moment(date).from(dateNow)}</time>
      </OverlayTrigger>
    )}
  </Consumer>
);
export default Date;
