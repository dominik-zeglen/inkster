/* tslint:disable */
// This file was automatically generated and should not be edited.

import { PaginationInput } from "./../../../types/globalTypes";

// ====================================================
// GraphQL query operation: RootDirectories
// ====================================================

export interface RootDirectories_getRootDirectories_edges_node {
  __typename: "Directory";
  id: string;
  name: string;
  isPublished: boolean;
}

export interface RootDirectories_getRootDirectories_edges {
  __typename: "DirectoryConnectionEdge";
  node: RootDirectories_getRootDirectories_edges_node;
}

export interface RootDirectories_getRootDirectories_pageInfo {
  __typename: "PageInfo";
  startCursor: string | null;
  endCursor: string | null;
  hasPreviousPage: boolean;
  hasNextPage: boolean;
}

export interface RootDirectories_getRootDirectories {
  __typename: "DirectoryConnection";
  edges: RootDirectories_getRootDirectories_edges[];
  pageInfo: RootDirectories_getRootDirectories_pageInfo;
}

export interface RootDirectories {
  getRootDirectories: RootDirectories_getRootDirectories | null;
}

export interface RootDirectoriesVariables {
  paginate: PaginationInput;
}
