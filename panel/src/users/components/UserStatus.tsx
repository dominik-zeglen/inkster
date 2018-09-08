import * as React from "react";
import { Checkbox, Panel } from "react-bootstrap";
import withStyles from "react-jss";

import i18n from "../../i18n";
import Date from "../../components/Date";
import Skeleton from "../../components/Skeleton";

interface Props {
  createdAt?: string;
  data: {
    isActive: boolean;
  };
  updatedAt?: string;
  onChange: (event: React.ChangeEvent<any>) => void;
}

const handleCheckboxChange = (event: React.ChangeEvent<any>) => {
  event.target.value = !event.target.value;
  return event;
};

const decorate = withStyles(
  (theme: any) => ({
    label: {
      marginRight: theme.spacing
    }
  }),
  { displayName: "UserStatus" }
);
export const UserStatus = decorate<Props>(
  ({ classes, createdAt, data, updatedAt, onChange }) => (
    <Panel>
      <Panel.Heading>
        <Panel.Title>{i18n.t("Status")}</Panel.Title>
      </Panel.Heading>
      <Panel.Body>
        <p>
          {createdAt ? (
            <>
              <span className={classes.label}>{i18n.t("Created")}</span>
              <Date date={createdAt} />
            </>
          ) : (
            <Skeleton />
          )}
          <br />
          {updatedAt ? (
            <>
              <span className={classes.label}>{i18n.t("Last modified")}</span>
              <Date date={updatedAt} />
            </>
          ) : (
            <Skeleton />
          )}
          <br />
        </p>
        <p>
          <Checkbox
            checked={data.isActive}
            name="isActive"
            onClick={handleCheckboxChange}
          >
            {i18n.t("Active")}
          </Checkbox>
        </p>
      </Panel.Body>
    </Panel>
  )
);
export default UserStatus;
