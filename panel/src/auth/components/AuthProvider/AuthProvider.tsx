import * as React from "react";
import { Mutation } from "react-apollo";

import {
  getAuthToken,
  removeAuthToken,
  setAuthToken,
  User,
  UserContext
} from ".";
import mTokenVerify from "../../queries/mTokenVerify";
import mLogin from "../../queries/mLogin";

interface AuthProviderOperationsProps {
  children:
    | ((
        props: {
          hasToken: boolean;
          isAuthenticated: boolean;
          tokenAuthLoading: boolean;
          tokenVerifyLoading: boolean;
        }
      ) => React.ReactElement<any>)
    | React.ReactNode;
  onError?: () => void;
}
const AuthProviderOperations: React.StatelessComponent<
  AuthProviderOperationsProps
> = ({ children, onError }) => {
  return (
    <Mutation mutation={mLogin} onError={onError}>
      {(login, loginData) => (
        <Mutation mutation={mTokenVerify} onError={onError}>
          {(tokenVerify, tokenVerifyData) => (
            <AuthProvider
              tokenAuth={{ ...loginData, mutate: login }}
              tokenVerify={{ ...tokenVerifyData, mutate: tokenVerify }}
            >
              {children}
            </AuthProvider>
          )}
        </Mutation>
      )}
    </Mutation>
  );
};

interface AuthProviderProps {
  children:
    | ((
        props: {
          hasToken: boolean;
          isAuthenticated: boolean;
          tokenAuthLoading: boolean;
          tokenVerifyLoading: boolean;
        }
      ) => React.ReactElement<any>)
    | React.ReactNode;
  tokenAuth: any;
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
    const { tokenAuth, tokenVerify } = props;
    if (tokenAuth.error || tokenVerify.error) {
      this.logout();
    }
    if (tokenAuth.data) {
      const user = tokenAuth.data.login.user;
      // FIXME: Now we set state also when auth fails and returned user is
      // `null`, because the LoginView uses this `null` to display error.
      this.setState({ user });
      if (user) {
        setAuthToken(tokenAuth.data.login.token, this.state.persistToken);
      }
    }
    if (tokenVerify.data && tokenVerify.data.verifyToken.user) {
      const user = tokenVerify.data.verifyToken.user;
      this.setState({ user });
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
    const { tokenAuth } = this.props;
    this.setState({ persistToken });
    tokenAuth.mutate({ variables: { email, password } });
  };

  logout = () => {
    this.setState({ user: undefined });
    removeAuthToken();
  };

  render() {
    const { children, tokenAuth, tokenVerify } = this.props;
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
              tokenAuthLoading: tokenAuth.loading,
              tokenVerifyLoading: tokenVerify.loading
            })
          : children}
      </UserContext.Provider>
    );
  }
}

export default AuthProviderOperations;
