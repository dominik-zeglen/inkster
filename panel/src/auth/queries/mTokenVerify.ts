import gql from "graphql-tag";

import { TypedMutation } from "../../api";
import { TokenVerify, TokenVerifyVariables } from "./types/TokenVerify";

const mTokenVerify = gql`
  mutation TokenVerify($token: String!) {
    verifyToken(token: $token) {
      result
      user {
        id
        email
      }
    }
  }
`;
export default TypedMutation<TokenVerify, TokenVerifyVariables>(mTokenVerify);
