import * as React from "react";
import { Panel } from "react-bootstrap";
import withStyles from "react-jss";

import i18n from "../../i18n";
import Date from "../../components/Date";
import Skeleton from "../../components/Skeleton";

interface Props {
  createdAt?: string;
  updatedAt?: string;
}

const decorate = withStyles(
  (theme: any) => ({
    label: {
      marginRight: theme.spacing
    }
  }),
  { displayName: "DirectoryStatus" }
);
export const DirectoryStatus = decorate<Props>(
  ({ classes, createdAt, updatedAt }) => (
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
      </Panel.Body>
    </Panel>
  )
);
export default DirectoryStatus;
