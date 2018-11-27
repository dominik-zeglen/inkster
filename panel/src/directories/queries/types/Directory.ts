/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: Directory
// ====================================================

export interface Directory_getDirectory_parent {
  __typename: "Directory";
  id: string;
}

export interface Directory_getDirectory_pages {
  __typename: "Page";
  id: string;
  name: string;
}

export interface Directory_getDirectory {
  __typename: "Directory";
  id: string;
  createdAt: string;
  updatedAt: string;
  name: string;
  isPublished: boolean;
  parent: Directory_getDirectory_parent | null;
  pages: (Directory_getDirectory_pages | null)[] | null;
}

export interface Directory {
  getDirectory: Directory_getDirectory | null;
}

export interface DirectoryVariables {
  id: string;
}
