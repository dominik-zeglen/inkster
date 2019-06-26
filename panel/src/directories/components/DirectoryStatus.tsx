import * as React from "react";
import CardHeader from "aurora-ui-kit/dist/components/CardHeader";
import Card from "aurora-ui-kit/dist/components/Card";
import CardTitle from "aurora-ui-kit/dist/components/CardTitle";
import CardContent from "aurora-ui-kit/dist/components/CardContent";
import Checkbox from "aurora-ui-kit/dist/components/Checkbox";

import i18n from "../../i18n";
import Date from "../../components/Date";
import { ITheme } from "aurora-ui-kit/dist/theme";
import createUseStyles from "aurora-ui-kit/dist/utils/jss";
import Skeleton from "aurora-ui-kit/dist/components/Skeleton";

interface Props {
  createdAt?: string;
  updatedAt?: string;
  data: {
    isPublished: boolean;
  };
  onChange: (event: React.ChangeEvent) => void;
}

const useStyles = createUseStyles((theme: ITheme) => ({
  label: {
    marginRight: theme.spacing,
  },
}));
export const DirectoryStatus: React.FC<Props> = ({
  createdAt,
  data,
  updatedAt,
  onChange,
}) => {
  const classes = useStyles();
  return (
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
          checked={data.isPublished}
          onChange={value =>
            onChange({
              target: {
                name: "isPublished",
                value,
              },
            } as any)
          }
        />
      </CardContent>
    </Card>
  );
};
export default DirectoryStatus;
