/* tslint:disable */
// This file was automatically generated and should not be edited.

import { UserUpdateInput } from "./../../../types/globalTypes";

// ====================================================
// GraphQL mutation operation: UpdateUser
// ====================================================

export interface UpdateUser_updateUser_errors {
  __typename: "InputError";
  field: string;
  message: string;
}

export interface UpdateUser_updateUser_user {
  __typename: "User";
  id: string;
  email: string;
  updatedAt: string;
  isActive: boolean;
}

export interface UpdateUser_updateUser {
  __typename: "UserOperationResult";
  errors: UpdateUser_updateUser_errors[];
  user: UpdateUser_updateUser_user | null;
}

export interface UpdateUser {
  updateUser: UpdateUser_updateUser;
}

export interface UpdateUserVariables {
  id: string;
  input: UserUpdateInput;
}
