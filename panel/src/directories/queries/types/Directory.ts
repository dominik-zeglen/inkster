/* tslint:disable */
// This file was automatically generated and should not be edited.

import { PaginationInput } from "./../../../types/globalTypes";

// ====================================================
// GraphQL query operation: Directory
// ====================================================

export interface Directory_getDirectory_parent {
  __typename: "Directory";
  id: string;
}

export interface Directory_getDirectory_pages_edges_node {
  __typename: "Page";
  id: string;
  name: string;
}

export interface Directory_getDirectory_pages_edges {
  __typename: "PageConnectionEdge";
  node: Directory_getDirectory_pages_edges_node;
}

export interface Directory_getDirectory_pages_pageInfo {
  __typename: "PageInfo";
  startCursor: string | null;
  endCursor: string | null;
  hasPreviousPage: boolean;
  hasNextPage: boolean;
}

export interface Directory_getDirectory_pages {
  __typename: "PageConnection";
  edges: Directory_getDirectory_pages_edges[];
  pageInfo: Directory_getDirectory_pages_pageInfo;
}

export interface Directory_getDirectory {
  __typename: "Directory";
  id: string;
  createdAt: string;
  updatedAt: string;
  name: string;
  isPublished: boolean;
  parent: Directory_getDirectory_parent | null;
  pages: Directory_getDirectory_pages | null;
}

export interface Directory {
  getDirectory: Directory_getDirectory | null;
}

export interface DirectoryVariables {
  id: string;
  paginate: PaginationInput;
}
