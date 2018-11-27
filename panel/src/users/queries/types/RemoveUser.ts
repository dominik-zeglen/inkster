/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: RemoveUser
// ====================================================

export interface RemoveUser_removeUser {
  __typename: "UserRemoveResult";
  removedObjectId: string | null;
}

export interface RemoveUser {
  removeUser: RemoveUser_removeUser;
}

export interface RemoveUserVariables {
  id: string;
}
