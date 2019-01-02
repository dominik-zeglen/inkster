/* tslint:disable */
// This file was automatically generated and should not be edited.

import { PaginationInput } from "./../../../types/globalTypes";

// ====================================================
// GraphQL query operation: UserList
// ====================================================

export interface UserList_users_edges_node {
  __typename: "User";
  id: string;
  email: string;
  isActive: boolean;
}

export interface UserList_users_edges {
  __typename: "UserConnectionEdge";
  node: UserList_users_edges_node;
}

export interface UserList_users_pageInfo {
  __typename: "PageInfo";
  startCursor: string | null;
  endCursor: string | null;
  hasPreviousPage: boolean;
  hasNextPage: boolean;
}

export interface UserList_users {
  __typename: "UserConnection";
  edges: UserList_users_edges[];
  pageInfo: UserList_users_pageInfo;
}

export interface UserList {
  users: UserList_users | null;
}

export interface UserListVariables {
  paginate: PaginationInput;
}
