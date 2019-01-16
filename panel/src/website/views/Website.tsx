import * as React from "react";

import Notificator from "../../components/Notificator";
import WebsitePage from "../components/WebsitePage";
import Website from "../queries/qWebsite";
import WebsiteUpdateMutation from "../queries/mWebsiteUpdate";
import { WebsiteUpdate } from "../queries/types/WebsiteUpdate";
import i18n from "../../i18n";

export const WebsiteView = () => (
  <Notificator>
    {notify => {
      const handleWebsiteUpdate = (data: WebsiteUpdate) => {
        if (data.updateWebsite.errors.length === 0) {
          notify({
            text: i18n.t("Website settings updated", {
              context: "notification",
            }),
          });
        }
      };
      return (
        <WebsiteUpdateMutation onCompleted={handleWebsiteUpdate}>
          {(websiteUpdate, websiteUpdateOpts) => (
            <Website>
              {({ data, loading }) => (
                <WebsitePage
                  disabled={loading}
                  transaction={
                    websiteUpdateOpts.loading ? "loading" : "default"
                  }
                  website={data.website}
                  onSubmit={formData =>
                    websiteUpdate({
                      variables: {
                        input: formData,
                      },
                    })
                  }
                />
              )}
            </Website>
          )}
        </WebsiteUpdateMutation>
      );
    }}
  </Notificator>
);
export default WebsiteView;
