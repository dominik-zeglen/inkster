import * as React from "react";
import withStyles from "react-jss";
import Input from "aurora-ui-kit/dist/components/TextInput";
import CardHeader from "aurora-ui-kit/dist/components/CardHeader";
import Card from "aurora-ui-kit/dist/components/Card";
import CardContent from "aurora-ui-kit/dist/components/CardContent";
import CardTitle from "aurora-ui-kit/dist/components/CardTitle";

import Container from "../../components/Container";
import Form from "../../components/Form";
import FormSave from "../../components/FormSave";
import PageHeader from "../../components/PageHeader";
import { Website_website } from "../queries/types/Website";
import i18n from "../../i18n";
import { maybe } from "../../utils";
import { TransactionState } from "../..";

interface FormData {
  name: string;
  description: string;
  domain: string;
}
interface Props {
  disabled: boolean;
  website: Website_website;
  transaction: TransactionState;
  onSubmit: (data: FormData) => void;
}

const decorate = withStyles(
  theme => ({
    container: {
      marginBottom: theme.spacing * 2,
    },
    root: {
      display: "grid" as "grid",
      gridColumnGap: theme.spacing + "px",
      gridTemplateColumns: "2fr 1fr",
    },
  }),
  { displayName: "UserDetailsPage" },
);
export const WebsitePage = decorate<Props>(
  ({ classes, disabled, onSubmit, transaction, website }) => {
    const initialForm: FormData = {
      description: maybe(() => website.description),
      domain: maybe(() => website.domain),
      name: maybe(() => website.name),
    };
    return (
      <Form
        initial={initialForm}
        onSubmit={onSubmit}
        key={JSON.stringify(website)}
      >
        {({ change, data, hasChanged, submit }) => (
          <Container width="md">
            <PageHeader title={i18n.t("Website Settings")} />
            <div className={classes.root}>
              <div>
                <div className={classes.container}>
                  <Card>
                    <CardHeader>
                      <CardTitle>{i18n.t("Basic settings")}</CardTitle>
                    </CardHeader>
                    <CardContent>
                      <div className={classes.container}>
                        <Input
                          disabled={disabled}
                          value={data.name}
                          label={i18n.t("Website name")}
                          onChange={value =>
                            change({
                              target: {
                                name: "name",
                                value,
                              },
                            } as any)
                          }
                        />
                      </div>
                      <div className={classes.container}>
                        <Input
                          disabled={disabled}
                          value={data.description}
                          label={i18n.t("Website description")}
                          onChange={value =>
                            change({
                              target: {
                                name: "description",
                                value,
                              },
                            } as any)
                          }
                        />
                      </div>
                    </CardContent>
                  </Card>
                </div>
                <div className={classes.container}>
                  <Card>
                    <CardHeader>
                      <CardTitle>{i18n.t("Advanced settings")}</CardTitle>
                    </CardHeader>
                    <CardContent>
                      <div className={classes.container}>
                        <Input
                          disabled={disabled}
                          value={data.domain}
                          label={i18n.t("Website domain")}
                          InputProps={{
                            componentProps: {
                              type: "url",
                            },
                          }}
                          onChange={value =>
                            change({
                              target: {
                                name: "domain",
                                value,
                              },
                            } as any)
                          }
                        />
                      </div>
                    </CardContent>
                  </Card>
                </div>
              </div>
            </div>
            <FormSave
              disabled={disabled || !hasChanged}
              variant={transaction}
              onConfirm={submit}
            />
          </Container>
        )}
      </Form>
    );
  },
);
export default WebsitePage;
