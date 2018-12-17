import * as React from "react";

import RootDirectories from "../queries/qRootDirectories";
import DirectoryCreateMutation from "../queries/mDirectoryCreate";
import DirectoryRootPage from "../components/DirectoryRootPage";
import Navigator from "../../components/Navigator";
import Notificator from "../../components/Notificator";
import urls from "../../urls";
import i18n from "../../i18n";
import {
  DirectoryCreateVariables,
  DirectoryCreate,
} from "../queries/types/DirectoryCreate";
import Input from "../../components/Input";
import FormDialog from "../../components/FormDialog";
import { mergeQs } from "../../utils";
import { Modal } from "../../types";

const dummy = () => {};

export type QueryParams = Partial<Modal<"create-directory">>;
export interface Props {
  params: QueryParams;
}

export const DirectoryRoot: React.StatelessComponent<Props> = ({ params }) => (
  <Notificator>
    {notify => (
      <Navigator>
        {navigate => {
          const handleRowClick = (id: string) => () =>
            navigate(urls.directoryDetails(id));
          const handleCreate = (data: DirectoryCreate) => {
            notify({
              text: i18n.t("Created directory", {
                context: "notification",
              }),
            });
            if (data.createDirectory.errors.length === 0) {
              navigate(
                urls.directoryDetails(data.createDirectory.directory.id),
              );
            }
          };
          return (
            <RootDirectories>
              {({ data, loading }) => (
                <DirectoryCreateMutation onCompleted={handleCreate}>
                  {addDirectory => {
                    const handleAddDirectory = (
                      variables: DirectoryCreateVariables,
                    ) => addDirectory({ variables });
                    return (
                      <>
                        <DirectoryRootPage
                          directories={
                            data ? data.getRootDirectories : undefined
                          }
                          disabled={loading}
                          loading={loading}
                          onAdd={() =>
                            navigate(
                              mergeQs(params, {
                                modal: "create-directory",
                              }),
                            )
                          }
                          onNextPage={dummy}
                          onPreviousPage={dummy}
                          onRowClick={handleRowClick}
                        />
                        <FormDialog
                          onClose={() =>
                            navigate(
                              mergeQs(params, {
                                modal: undefined,
                              }),
                            )
                          }
                          onConfirm={handleAddDirectory}
                          show={params.modal === "create-directory"}
                          title={i18n.t("Create directory")}
                          width="sm"
                          initial={{ name: "" }}
                        >
                          {({ change, data: formData }) => (
                            <Input
                              name="name"
                              onChange={change}
                              value={formData.name}
                              label={i18n.t("Directory name")}
                            />
                          )}
                        </FormDialog>
                      </>
                    );
                  }}
                </DirectoryCreateMutation>
              )}
            </RootDirectories>
          );
        }}
      </Navigator>
    )}
  </Notificator>
);
export default DirectoryRoot;
