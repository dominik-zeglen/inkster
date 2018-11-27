/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: UserDetails
// ====================================================

export interface UserDetails_user {
  __typename: "User";
  id: string;
  email: string;
  isActive: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface UserDetails {
  user: UserDetails_user | null;
}

export interface UserDetailsVariables {
  id: string;
}
