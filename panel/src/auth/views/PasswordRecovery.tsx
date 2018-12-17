import * as React from "react";
import { RouteComponentProps } from "react-router-dom";
import { parse as parseQs } from "qs";

import PasswordResetPage, {
  FormData as PasswordResetPageFormData,
} from "../components/PasswordResetPage";
import PasswordResetSendEmailPage, {
  FormData as PasswordResetSendEmailPageFormData,
} from "../components/PasswordResetSendEmailPage";
import Navigator from "../../components/Navigator";
import Notificator, { NotificationType } from "../../components/Notificator";
import i18n from "../../i18n";
import ResetUserPasswordMutation from "../queries/mResetUserPassword";
import SendUserPasswordResetTokenMutation from "../queries/mSendUserPasswordResetToken";
import { ResetUserPassword } from "../queries/types/ResetUserPassword";
import { SendUserPasswordResetToken } from "../queries/types/SendUserPasswordResetToken";
import urls from "../../urls";

const PasswordRecoveryView: React.StatelessComponent<
  RouteComponentProps<{}>
> = ({ location }) => {
  const token = parseQs(location.search.substr(1)).token;
  return (
    <Navigator>
      {navigate => (
        <Notificator>
          {notify => {
            const handlePasswordResetComplete = (data: ResetUserPassword) => {
              if (data.resetUserPassword) {
                notify({
                  text: i18n.t(
                    "Password changed, you can now log in to your account",
                  ),
                });
              } else {
                notify({
                  text: i18n.t("Invalid token"),
                  type: NotificationType.ERROR,
                });
              }
              navigate(urls.home);
            };
            const handleEmailSendComplete = (
              data: SendUserPasswordResetToken,
            ) => {
              if (data.sendUserPasswordResetToken) {
                notify({
                  text: i18n.t(
                    "An e-mail containing link to password reset had been sent",
                  ),
                });
                navigate(urls.home);
              } else {
                notify({
                  text: i18n.t("Could not send e-mail"),
                  type: NotificationType.ERROR,
                });
              }
            };

            return (
              <SendUserPasswordResetTokenMutation
                onCompleted={handleEmailSendComplete}
              >
                {(sendEmail, sendEmailOpts) => (
                  <ResetUserPasswordMutation
                    onCompleted={handlePasswordResetComplete}
                  >
                    {(resetPassword, resetPasswordOpts) => {
                      const handlePasswordReset = (
                        data: PasswordResetPageFormData,
                      ) =>
                        resetPassword({
                          variables: {
                            token,
                            password: data.password,
                          },
                        });
                      const handleEmailSend = (
                        data: PasswordResetSendEmailPageFormData,
                      ) => sendEmail({ variables: data });
                      return token ? (
                        <PasswordResetPage
                          disabled={resetPasswordOpts.loading}
                          onSubmit={handlePasswordReset}
                        />
                      ) : (
                        <PasswordResetSendEmailPage
                          disabled={sendEmailOpts.loading}
                          onSubmit={handleEmailSend}
                        />
                      );
                    }}
                  </ResetUserPasswordMutation>
                )}
              </SendUserPasswordResetTokenMutation>
            );
          }}
        </Notificator>
      )}
    </Navigator>
  );
};
export default PasswordRecoveryView;
