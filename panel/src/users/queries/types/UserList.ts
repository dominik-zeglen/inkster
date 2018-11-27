/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: UserList
// ====================================================

export interface UserList_users {
  __typename: "User";
  id: string;
  email: string;
  isActive: boolean;
}

export interface UserList {
  users: (UserList_users | null)[] | null;
}
