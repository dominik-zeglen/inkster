import * as React from "react";
import { Mutation } from "react-apollo";
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
import mResetUserPassword, {
  Result as mResetUserPasswordResult,
} from "../queries/mResetUserPassword";
import mSendUserPasswordResetToken, {
  Result as mSendUserPasswordResetTokenResult,
} from "../queries/mSendUserPasswordResetToken";

const PasswordRecoveryView: React.StatelessComponent<
  RouteComponentProps<{}>
> = ({ location }) => {
  const token = parseQs(location.search.substr(1)).token;
  return (
    <Navigator>
      {navigate => (
        <Notificator>
          {notify => {
            const handlePasswordResetComplete = (
              data: mResetUserPasswordResult,
            ) => {
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
              navigate("/");
            };
            const handleEmailSendComplete = (
              data: mSendUserPasswordResetTokenResult,
            ) => {
              if (data.sendUserPasswordResetToken) {
                notify({
                  text: i18n.t(
                    "An e-mail containing link to password reset had been sent",
                  ),
                });
                navigate("/");
              } else {
                notify({
                  text: i18n.t("Could not send e-mail"),
                  type: NotificationType.ERROR,
                });
              }
            };

            return (
              <Mutation
                mutation={mSendUserPasswordResetToken}
                onCompleted={handleEmailSendComplete}
              >
                {(sendEmail, sendEmailOpts) => (
                  <Mutation
                    mutation={mResetUserPassword}
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
                  </Mutation>
                )}
              </Mutation>
            );
          }}
        </Notificator>
      )}
    </Navigator>
  );
};
export default PasswordRecoveryView;
