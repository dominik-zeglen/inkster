import * as React from "react";
import { Query } from "react-apollo";
import { ApolloError } from "apollo-client";

import qPage from "../../queries/qPage";
import PageDetailsPage, { FormData } from "../../components/PageDetailsPage";
import Navigator from "../../../components/Navigator";
import Notificator, {NotificationType} from "../../../components/Notificator";
import MutationProvider from "./MutationProvider";
import { urls } from "../../../";
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
    transaction: "default" as "default"
  };

  handleUpdate = () => {
    this.setState({ transaction: "success" });
    setTimeout(() => this.setState({ transaction: "default" }), 3000);
  };
  handleUpdateError = (event: ApolloError) => {
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
              <Query
                query={qPage}
                variables={{ id }}
                fetchPolicy="network-only"
              >
                {({ data, error, loading }) => {
                  if (error) {
                    console.error(error);
                    return JSON.stringify(error);
                  }
                  const handleBack = () =>
                    navigate(
                      urls.directoryDetails(
                        data && data.page && data.page.parent
                          ? data.page.parent.id
                          : undefined
                      )
                    );
                  const handleDelete = () => {
                    notify({
                      text: i18n.t("Page deleted", {
                        context: "notification"
                      })
                    });
                    handleBack();
                  };
                  const handleError = () => notify({
                      text: i18n.t("Something has gone wrong", {
                        context: "notification"
                      }),
                      type: NotificationType.ERROR
                    });
                  return (
                    <WithUpload>
                      {uploadFile => {
                        const handleUpload = (onChange: any) => (
                          event: React.ChangeEvent<any>
                        ) => {
                          uploadFile.uploadFile({
                            file: event.target.files[0],
                            onSuccess: filename =>
                              onChange({
                                target: {
                                  name: "value",
                                  value: filename
                                }
                              } as any),
                            onError: handleError
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
                              const formLoading = updatePage.loading;
                              const modalLoading = deletePage.loading;

                              const handleSubmit = (formData: FormData) =>
                                updatePage.mutate({
                                  id,
                                  input: {
                                    isPublished: formData.isPublished,
                                    name: formData.name,
                                    slug: formData.slug,
                                    fields: formData.fields.map(f => ({
                                      name: f.id,
                                      update: { name: f.name, value: f.value }
                                    }))
                                  },
                                  add: formData.addFields,
                                  remove: formData.removeFields
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
                                  onDelete={deletePage.mutate}
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
              </Query>
            )}
          </Navigator>
        )}
      </Notificator>
    );
  }
}
export default PageDetails;
