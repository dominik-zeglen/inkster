/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: Viewer
// ====================================================

export interface Viewer_viewer_pages {
  __typename: "Page";
  id: string;
  name: string;
  slug: string;
  isPublished: boolean;
}

export interface Viewer_viewer {
  __typename: "User";
  id: string;
  email: string;
  pages: Viewer_viewer_pages[];
}

export interface Viewer {
  viewer: Viewer_viewer | null;
}
