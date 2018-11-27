/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: TokenVerify
// ====================================================

export interface TokenVerify_verifyToken_user {
  __typename: "User";
  id: string;
  email: string;
}

export interface TokenVerify_verifyToken {
  __typename: "VerifyTokenResult";
  result: boolean;
  user: TokenVerify_verifyToken_user | null;
}

export interface TokenVerify {
  verifyToken: TokenVerify_verifyToken | null;
}

export interface TokenVerifyVariables {
  token: string;
}
