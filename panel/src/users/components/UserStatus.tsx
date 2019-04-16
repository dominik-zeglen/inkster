import * as React from "react";
import Card from "aurora-ui-kit/dist/components/Card";
import CardContent from "aurora-ui-kit/dist/components/CardContent";
import CardHeader from "aurora-ui-kit/dist/components/CardHeader";
import CardTitle from "aurora-ui-kit/dist/components/CardTitle";
import withStyles from "react-jss";
import Checkbox from "aurora-ui-kit/dist/components/Checkbox";

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

const decorate = withStyles(
  (theme: any) => ({
    label: {
      marginRight: theme.spacing,
    },
  }),
  { displayName: "UserStatus" },
);
export const UserStatus = decorate<Props>(
  ({ classes, createdAt, data, updatedAt, onChange }) => (
    <Card>
      <CardHeader>
        <CardTitle>{i18n.t("Status")}</CardTitle>
      </CardHeader>
      <CardContent>
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
        <Checkbox
          checked={data.isActive}
          label={i18n.t("Active")}
          onChange={value =>
            onChange({
              target: {
                name: "isActive",
                value,
              },
            } as any)
          }
        />
      </CardContent>
    </Card>
  ),
);
export default UserStatus;
