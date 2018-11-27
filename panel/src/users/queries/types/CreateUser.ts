/* tslint:disable */
// This file was automatically generated and should not be edited.

import { UserCreateInput } from "./../../../types/globalTypes";

// ====================================================
// GraphQL mutation operation: CreateUser
// ====================================================

export interface CreateUser_createUser_errors {
  __typename: "InputError";
  field: string;
  message: string;
}

export interface CreateUser_createUser_user {
  __typename: "User";
  id: string;
  email: string;
  createdAt: string;
  updatedAt: string;
  isActive: boolean;
}

export interface CreateUser_createUser {
  __typename: "UserOperationResult";
  errors: CreateUser_createUser_errors[];
  user: CreateUser_createUser_user | null;
}

export interface CreateUser {
  createUser: CreateUser_createUser;
}

export interface CreateUserVariables {
  input: UserCreateInput;
}
