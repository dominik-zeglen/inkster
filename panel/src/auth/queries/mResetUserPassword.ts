import gql from "graphql-tag";

const mResetUserPassword = gql`
  mutation ResetUserPassword($token: String!, $password: String!) {
    resetUserPassword(token: $token, password: $password)
  }
`;
export interface Variables {
  token: string;
  password: string;
}
export interface Result {
  resetUserPassword: boolean;
}
export default mResetUserPassword;
