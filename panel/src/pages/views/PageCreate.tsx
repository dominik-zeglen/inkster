import * as React from "react";

import { TransactionState } from "../../";
import urls from "../../urls";
import Navigator from "../../components/Navigator";
import Notificator, { NotificationType } from "../../components/Notificator";
import PageCreatePage, { FormData } from "../components/PageCreatePage";
import PageCreateMutation from "../queries/mPageCreate";
import i18n from "../../i18n";
import { WithUpload } from "../../UploadProvider";
import { PageCreate } from "../queries/types/PageCreate";

interface Props {
  directory: string;
}
interface State {
  transaction: TransactionState;
}
export class PageCreateView extends React.Component<Props, State> {
  state = {
    transaction: "default" as "default",
  };

  handleSubmitSuccess = (cb: () => void) => {
    this.setState({ transaction: "success" });
    setTimeout(cb, 3000);
  };

  render() {
    const { directory } = this.props;
    return (
      <Notificator>
        {notify => (
          <Navigator>
            {navigate => {
              const handleBack = () =>
                navigate(urls.directoryDetails(directory));
              const handleSubmitCompleted = (data: PageCreate) => {
                if (data.createPage.errors.length > 0) {
                  notify({
                    text: i18n.t("Something has gone wrong", {
                      context: "notification",
                    }),
                    type: NotificationType.ERROR,
                  });
                }
                notify({
                  text: i18n.t("Page created", {
                    context: "notification",
                  }),
                });
                navigate(urls.pageDetails(data.createPage.page.id));
              };
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
                        onError: () => console.log("not ok"),
                      });
                    };
                    return (
                      <PageCreateMutation onCompleted={handleSubmitCompleted}>
                        {(createPage, { data, loading }) => {
                          const handleSubmit = (formData: FormData) =>
                            createPage({
                              variables: {
                                name: formData.name,
                                parentId: directory,
                                fields: formData.addFields,
                              },
                            });
                          return (
                            <PageCreatePage
                              disabled={loading}
                              loading={loading}
                              title={i18n.t("Create new page")}
                              transaction={this.state.transaction}
                              onBack={handleBack}
                              onUpload={handleUpload}
                              onSubmit={handleSubmit}
                            />
                          );
                        }}
                      </PageCreateMutation>
                    );
                  }}
                </WithUpload>
              );
            }}
          </Navigator>
        )}
      </Notificator>
    );
  }
}
export default PageCreateView;
