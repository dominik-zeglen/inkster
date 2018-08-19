import * as React from "react";
import * as moment from "moment";
import { OverlayTrigger, Tooltip } from "react-bootstrap";

import { Consumer } from "./DateContext";

interface Props {
  date: string;
  locale: string;
}

export const Date: React.StatelessComponent<Props> = ({ date, locale }) => (
  <Consumer>
    {dateNow => (
      <OverlayTrigger
        placement="bottom"
        overlay={
          <Tooltip>
            {moment(date)
              .toDate()
              .toLocaleString()}
          </Tooltip>
        }
      >
        {moment(date).from(dateNow)}
      </OverlayTrigger>
    )}
  </Consumer>
);
export default Date;
