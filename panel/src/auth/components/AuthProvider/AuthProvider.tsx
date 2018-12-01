import * as React from "react";

import {
  getAuthToken,
  removeAuthToken,
  setAuthToken,
  User,
  UserContext,
} from ".";
import TokenVerifyMutation from "../../queries/mTokenVerify";
import LoginMutation from "../../queries/mLogin";

interface AuthProviderOperationsProps {
  children:
    | ((
        props: {
          hasToken: boolean;
          isAuthenticated: boolean;
          loginLoading: boolean;
          tokenVerifyLoading: boolean;
        },
      ) => React.ReactElement<any>)
    | React.ReactNode;
  onError?: () => void;
}
const AuthProviderOperations: React.StatelessComponent<
  AuthProviderOperationsProps
> = ({ children, onError }) => {
  return (
    <LoginMutation onError={onError}>
      {(login, loginData) => (
        <TokenVerifyMutation onError={onError}>
          {(tokenVerify, tokenVerifyData) => (
            <AuthProvider
              login={{ ...loginData, mutate: login }}
              tokenVerify={{ ...tokenVerifyData, mutate: tokenVerify }}
            >
              {children}
            </AuthProvider>
          )}
        </TokenVerifyMutation>
      )}
    </LoginMutation>
  );
};

interface AuthProviderProps {
  children:
    | ((
        props: {
          hasToken: boolean;
          isAuthenticated: boolean;
          loginLoading: boolean;
          tokenVerifyLoading: boolean;
        },
      ) => React.ReactElement<any>)
    | React.ReactNode;
  login: any;
  tokenVerify: any;
}

interface AuthProviderState {
  user?: User;
  persistToken: boolean;
}

class AuthProvider extends React.Component<
  AuthProviderProps,
  AuthProviderState
> {
  constructor(props) {
    super(props);
    this.state = { user: undefined, persistToken: false };
  }

  componentWillReceiveProps(props: AuthProviderProps) {
    const { login, tokenVerify } = props;
    if (login.error || tokenVerify.error) {
      this.logout();
    }
    if (login.data) {
      const user = login.data.login.user;
      // FIXME: Now we set state also when auth fails and returned user is
      // `null`, because the LoginView uses this `null` to display error.
      this.setState({ user });
      if (user) {
        setAuthToken(login.data.login.token, this.state.persistToken);
      }
    } else {
      if (tokenVerify.data && tokenVerify.data.verifyToken.user) {
        const user = tokenVerify.data.verifyToken.user;
        this.setState({ user });
      }
    }
  }

  componentDidMount() {
    const { user } = this.state;
    const { tokenVerify } = this.props;
    const token = getAuthToken();
    if (!!token && !user) {
      tokenVerify.mutate({ variables: { token } });
    }
  }

  login = (email: string, password: string, persistToken: boolean) => {
    const { login } = this.props;
    this.setState({ persistToken });
    login.mutate({ variables: { email, password } });
  };

  logout = () => {
    this.setState({ user: undefined });
    removeAuthToken();
  };

  render() {
    const { children, login, tokenVerify } = this.props;
    const { user } = this.state;
    const isAuthenticated = !!user;
    return (
      <UserContext.Provider
        value={{ user, login: this.login, logout: this.logout }}
      >
        {typeof children === "function"
          ? children({
              hasToken: !!getAuthToken(),
              isAuthenticated,
              loginLoading: login.loading,
              tokenVerifyLoading: tokenVerify.loading,
            })
          : children}
      </UserContext.Provider>
    );
  }
}

export default AuthProviderOperations;
