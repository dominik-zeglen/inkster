import gql from "graphql-tag";

import { TypedMutation } from "../../api";
import {
  ResetUserPassword,
  ResetUserPasswordVariables,
} from "./types/ResetUserPassword";

const mResetUserPassword = gql`
  mutation ResetUserPassword($token: String!, $password: String!) {
    resetUserPassword(token: $token, password: $password)
  }
`;
export default TypedMutation<ResetUserPassword, ResetUserPasswordVariables>(
  mResetUserPassword,
);
