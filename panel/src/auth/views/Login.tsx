import * as React from "react";

import LoginPage, { FormData } from "../components/LoginPage";
import { UserContext } from "../components/AuthProvider";
import urls from '../../urls'

interface Props {
  loading: boolean;
}

const LoginView: React.StatelessComponent<Props> = ({ loading }) => (
  <UserContext.Consumer>
    {({ login, user }) => {
      const handleSubmit = (data: FormData) =>
        login(data.email, data.password, data.remember);
      return (
        <LoginPage
          disabled={loading}
          error={loading ? false : user === null}
          passwordRecoveryHref={urls.passwordRecovery}
          onSubmit={handleSubmit}
        />
      );
    }}
  </UserContext.Consumer>
);

export default LoginView;
