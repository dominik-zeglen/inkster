import * as React from "react";
import Card from "aurora-ui-kit/dist/components/Card";
import CardHeader from "aurora-ui-kit/dist/components/CardHeader";
import CardTitle from "aurora-ui-kit/dist/components/CardTitle";
import Checkbox from "aurora-ui-kit/dist/components/Checkbox";
import CardContent from "aurora-ui-kit/dist/components/CardContent";
import Skeleton from "aurora-ui-kit/dist/components/Skeleton";

import i18n from "../../i18n";
import Date from "../../components/Date";

interface Props {
  createdAt?: string;
  updatedAt?: string;
  data: {
    isPublished: boolean;
  };
  onChange: (event: React.ChangeEvent) => void;
}
export const PageStatus: React.FC<Props> = ({
  createdAt,
  data,
  updatedAt,
  onChange,
}) => (
  <Card>
    <CardHeader>
      <CardTitle>{i18n.t("Status")}</CardTitle>
    </CardHeader>
    <CardContent>
      <p>
        {createdAt ? (
          <>
            {i18n.t("Created")} <Date date={createdAt} />
          </>
        ) : (
          <Skeleton />
        )}
        <br />
        {updatedAt ? (
          <>
            {i18n.t("Last modified")} <Date date={updatedAt} />
          </>
        ) : (
          <Skeleton />
        )}
        <br />
      </p>
      <Checkbox
        checked={data.isPublished}
        label={i18n.t("Published")}
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
export default PageStatus;
