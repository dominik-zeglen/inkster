import * as React from "react";
import Input from "aurora-ui-kit/dist/components/TextInput";

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
import FormDialog from "../../components/FormDialog";
import { mergeQs, maybe } from "../../utils";
import { Modal, Pagination } from "../../types";
import Paginator, { createPaginationState } from "../../components/Paginator";
import { PAGINATE_BY } from "../..";

export type QueryParams = Partial<Modal<"create-directory"> & Pagination>;
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
          const paginationState = createPaginationState(PAGINATE_BY, params);
          return (
            <RootDirectories variables={{ paginate: paginationState }}>
              {({ data, loading }) => (
                <DirectoryCreateMutation onCompleted={handleCreate}>
                  {addDirectory => {
                    const handleAddDirectory = (
                      variables: DirectoryCreateVariables,
                    ) => addDirectory({ variables });
                    return (
                      <>
                        <Paginator
                          pageInfo={maybe(
                            () => data.getRootDirectories.pageInfo,
                          )}
                          paginationState={paginationState}
                          queryString={params}
                        >
                          {({ loadNextPage, loadPreviousPage, pageInfo }) => (
                            <DirectoryRootPage
                              directories={maybe(() =>
                                data.getRootDirectories.edges.map(
                                  edge => edge.node,
                                ),
                              )}
                              disabled={loading}
                              loading={loading}
                              onAdd={() =>
                                navigate(
                                  mergeQs(params, {
                                    modal: "create-directory",
                                  }),
                                )
                              }
                              onNextPage={loadNextPage}
                              onPreviousPage={loadPreviousPage}
                              onRowClick={handleRowClick}
                              pageInfo={pageInfo}
                            />
                          )}
                        </Paginator>
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
                          width="xs"
                          initial={{ name: "" }}
                        >
                          {({ change, data: formData }) => (
                            <Input
                              onChange={value =>
                                change({
                                  target: {
                                    name: "name",
                                    value,
                                  },
                                } as any)
                              }
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
