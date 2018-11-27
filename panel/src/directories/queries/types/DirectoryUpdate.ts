/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: DirectoryUpdate
// ====================================================

export interface DirectoryUpdate_updateDirectory_errors {
  __typename: "InputError";
  code: number;
  field: string;
}

export interface DirectoryUpdate_updateDirectory_directory {
  __typename: "Directory";
  id: string;
  updatedAt: string;
  name: string;
  isPublished: boolean;
}

export interface DirectoryUpdate_updateDirectory {
  __typename: "DirectoryOperationResult";
  errors: DirectoryUpdate_updateDirectory_errors[];
  directory: DirectoryUpdate_updateDirectory_directory | null;
}

export interface DirectoryUpdate {
  updateDirectory: DirectoryUpdate_updateDirectory | null;
}

export interface DirectoryUpdateVariables {
  id: string;
  name: string;
  isPublished?: boolean | null;
}
