/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: RootDirectories
// ====================================================

export interface RootDirectories_getRootDirectories {
  __typename: "Directory";
  id: string;
  name: string;
  isPublished: boolean;
}

export interface RootDirectories {
  getRootDirectories: (RootDirectories_getRootDirectories | null)[] | null;
}
