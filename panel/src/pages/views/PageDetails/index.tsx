import * as React from "react";

import Page from "../../queries/qPage";
import PageDetailsPage, { FormData } from "../../components/PageDetailsPage";
import Navigator from "../../../components/Navigator";
import Notificator, { NotificationType } from "../../../components/Notificator";
import MutationProvider from "./MutationProvider";
import urls from "../../../urls";
import i18n from "../../../i18n";
import { TransactionState } from "../../../";
import { WithUpload } from "../../../UploadProvider";

interface Props {
  id: string;
}
interface State {
  transaction: TransactionState;
}
export class PageDetails extends React.Component<Props, State> {
  state = {
    transaction: "default" as "default",
  };

  handleUpdate = () => {
    this.setState({ transaction: "success" });
    setTimeout(() => this.setState({ transaction: "default" }), 3000);
  };
  handleUpdateError = () => {
    this.setState({ transaction: "error" });
    setTimeout(() => this.setState({ transaction: "default" }), 3000);
  };

  render() {
    const { id } = this.props;
    return (
      <Notificator>
        {notify => (
          <Navigator>
            {navigate => (
              <Page variables={{ id }}>
                {({ data }) => {
                  const handleBack = () =>
                    navigate(
                      urls.directoryDetails(
                        data && data.page && data.page.parent
                          ? data.page.parent.id
                          : undefined,
                      ),
                    );
                  const handleDelete = () => {
                    notify({
                      text: i18n.t("Page deleted", {
                        context: "notification",
                      }),
                    });
                    handleBack();
                  };
                  const handleError = () =>
                    notify({
                      text: i18n.t("Something has gone wrong", {
                        context: "notification",
                      }),
                      type: NotificationType.ERROR,
                    });
                  return (
                    <WithUpload>
                      {uploadFile => {
                        const handleUpload = (onChange: any) => (
                          event: React.ChangeEvent<any>,
                        ) => {
                          uploadFile.uploadFile({
                            file: event.target.files[0],
                            onSuccess: filename =>
                              onChange({
                                target: {
                                  name: "value",
                                  value: filename,
                                },
                              } as any),
                            onError: handleError,
                          });
                        };
                        return (
                          <MutationProvider
                            id={id}
                            onPageUpdate={this.handleUpdate}
                            onPageDelete={handleDelete}
                            onError={handleError}
                          >
                            {({ deletePage, updatePage }) => {
                              const formLoading = updatePage.opts.loading;
                              const modalLoading = deletePage.opts.loading;

                              const handleSubmit = (formData: FormData) =>
                                updatePage.mutate({
                                  id,
                                  input: {
                                    isPublished: formData.isPublished,
                                    name: formData.name,
                                    slug: formData.slug,
                                  },
                                  add:
                                    formData.addFields.length > 0
                                      ? formData.addFields
                                      : null,
                                  remove:
                                    formData.removeFields.length > 0
                                      ? formData.removeFields
                                      : null,
                                  update: formData.fields.map(field => ({
                                    id: field.id,
                                    input: {
                                      name: field.name,
                                      value: field.value,
                                    },
                                  })),
                                });
                              return (
                                <PageDetailsPage
                                  disabled={formLoading || modalLoading}
                                  loading={formLoading || modalLoading}
                                  title={
                                    data && data.page
                                      ? data.page.name
                                      : undefined
                                  }
                                  transaction={this.state.transaction}
                                  page={data ? data.page : undefined}
                                  onBack={handleBack}
                                  onDelete={() => deletePage.mutate({ id })}
                                  onUpload={handleUpload}
                                  onSubmit={handleSubmit}
                                />
                              );
                            }}
                          </MutationProvider>
                        );
                      }}
                    </WithUpload>
                  );
                }}
              </Page>
            )}
          </Navigator>
        )}
      </Notificator>
    );
  }
}
export default PageDetails;
