/* tslint:disable */
// This file was automatically generated and should not be edited.

import { PaginationInput } from "./../../../types/globalTypes";

// ====================================================
// GraphQL query operation: UserDetails
// ====================================================

export interface UserDetails_user_pages_edges_node {
  __typename: "Page";
  id: string;
  name: string;
  createdAt: string;
  isPublished: boolean;
}

export interface UserDetails_user_pages_edges {
  __typename: "PageConnectionEdge";
  node: UserDetails_user_pages_edges_node;
}

export interface UserDetails_user_pages_pageInfo {
  __typename: "PageInfo";
  startCursor: string | null;
  endCursor: string | null;
  hasPreviousPage: boolean;
  hasNextPage: boolean;
}

export interface UserDetails_user_pages {
  __typename: "PageConnection";
  edges: UserDetails_user_pages_edges[];
  pageInfo: UserDetails_user_pages_pageInfo;
}

export interface UserDetails_user {
  __typename: "User";
  id: string;
  email: string;
  isActive: boolean;
  createdAt: string;
  updatedAt: string;
  pages: UserDetails_user_pages | null;
}

export interface UserDetails {
  user: UserDetails_user | null;
}

export interface UserDetailsVariables {
  id: string;
  paginate: PaginationInput;
}
