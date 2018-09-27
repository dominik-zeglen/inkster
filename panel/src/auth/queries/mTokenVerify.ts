import gql from "graphql-tag";

export interface Variables {
  token: string;
}
export interface Result {
  result: boolean;
  user?: {
    id: string;
    email: string;
  }
}

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
export default mTokenVerify;
