/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: Viewer
// ====================================================

export interface Viewer_viewer_pages_edges_node {
  __typename: "Page";
  id: string;
  name: string;
  slug: string;
  isPublished: boolean;
}

export interface Viewer_viewer_pages_edges {
  __typename: "PageConnectionEdge";
  node: Viewer_viewer_pages_edges_node;
}

export interface Viewer_viewer_pages_pageInfo {
  __typename: "PageInfo";
  startCursor: string | null;
  endCursor: string | null;
  hasPreviousPage: boolean;
  hasNextPage: boolean;
}

export interface Viewer_viewer_pages {
  __typename: "PageConnection";
  edges: Viewer_viewer_pages_edges[];
  pageInfo: Viewer_viewer_pages_pageInfo;
}

export interface Viewer_viewer {
  __typename: "User";
  id: string;
  email: string;
  pages: Viewer_viewer_pages | null;
}

export interface Viewer {
  viewer: Viewer_viewer | null;
}
