import * as React from "react";
import withStyles from "react-jss";

import Container from "../../components/Container";
import Form from "../../components/Form";
import FormSave from "../../components/FormSave";
import PageHeader from "../../components/PageHeader";
import { Website_website } from "../queries/types/Website";
import i18n from "../../i18n";
import { maybe } from "../../utils";
import { Panel } from "react-bootstrap";
import Input from "../../components/Input";
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
                  <Panel>
                    <Panel.Heading>
                      <Panel.Title>{i18n.t("Basic settings")}</Panel.Title>
                    </Panel.Heading>
                    <Panel.Body>
                      <div className={classes.container}>
                        <Input
                          disabled={disabled}
                          name={"name" as keyof FormData}
                          value={data.name}
                          label={i18n.t("Website name")}
                          onChange={change}
                        />
                      </div>
                      <div className={classes.container}>
                        <Input
                          disabled={disabled}
                          name={"description" as keyof FormData}
                          value={data.description}
                          label={i18n.t("Website description")}
                          onChange={change}
                        />
                      </div>
                    </Panel.Body>
                  </Panel>
                </div>
                <div className={classes.container}>
                  <Panel>
                    <Panel.Heading>
                      <Panel.Title>{i18n.t("Advanced settings")}</Panel.Title>
                    </Panel.Heading>
                    <Panel.Body>
                      <div className={classes.container}>
                        <Input
                          disabled={disabled}
                          name={"domain" as keyof FormData}
                          value={data.domain}
                          label={i18n.t("Website domain")}
                          type="url"
                          onChange={change}
                        />
                      </div>
                    </Panel.Body>
                  </Panel>
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
