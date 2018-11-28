import gql from "graphql-tag";

import { TypedMutation } from "../../api";
import { Login, LoginVariables } from "./types/Login";

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
export default TypedMutation<Login, LoginVariables>(mLogin);
