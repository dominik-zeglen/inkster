import gql from "graphql-tag";

export interface Variables {
  email: string;
  password: string;
}
export interface Result {
  token?: string;
  user?: {
    id: string;
    email: string;
  };
}

export const mLogin = gql`
  mutation Login($email: String!, $password: String!) {
    login(email: $email, password: $password) {
      token
      user {
        id
        email
      }
    }
  }
`;
export default mLogin;
