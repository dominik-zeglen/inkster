import * as React from "react";
import withStyles from "react-jss";
import Card from "aurora-ui-kit/dist/components/Card";
import CardHeader from "aurora-ui-kit/dist/components/CardHeader";
import CardTitle from "aurora-ui-kit/dist/components/CardTitle";
import CardContent from "aurora-ui-kit/dist/components/CardContent";

import i18n from "../../i18n";
import Checkbox from "../../components/Checkbox";
import Date from "../../components/Date";
import Skeleton from "../../components/Skeleton";

interface Props {
  createdAt?: string;
  updatedAt?: string;
  data: {
    isPublished: boolean;
  };
  onChange: (event: React.ChangeEvent) => void;
}

const decorate = withStyles(
  theme => ({
    label: {
      marginRight: theme.spacing,
    },
  }),
  { displayName: "PageStatus" },
);
export const PageStatus = decorate<Props>(
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
          label={i18n.t("Published")}
          name="isPublished"
          value={data.isPublished}
          onChange={onChange}
        />
      </CardContent>
    </Card>
  ),
);
export default PageStatus;
