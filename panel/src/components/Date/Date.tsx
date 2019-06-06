import * as React from "react";
import * as moment from "moment";
import Popover from "aurora-ui-kit/dist/components/Popover";

import DateContext from "./DateContext";

interface Props {
  date: string;
}

export const Date: React.StatelessComponent<Props> = ({ date }) => {
  const dateNow = React.useContext(DateContext);
  const anchor = React.useRef<HTMLDivElement>(null);
  const [hover, setHover] = React.useState(false);

  return (
    <Popover
      content={moment(date)
        .toDate()
        .toLocaleString()}
      isOpen={hover}
      enterExitTransitionDurationMs={0}
      place="below"
    >
      <time
        ref={anchor}
        onMouseEnter={() => setHover(true)}
        onMouseLeave={() => setHover(false)}
      >
        {moment(date).from(dateNow)}
      </time>
    </Popover>
  );
};
export default Date;
