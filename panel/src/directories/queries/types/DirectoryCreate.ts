/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: DirectoryCreate
// ====================================================

export interface DirectoryCreate_createDirectory_errors {
  __typename: "InputError";
  field: string;
  code: number;
}

export interface DirectoryCreate_createDirectory_directory_parent {
  __typename: "Directory";
  id: string;
}

export interface DirectoryCreate_createDirectory_directory {
  __typename: "Directory";
  id: string;
  createdAt: string;
  updatedAt: string;
  name: string;
  isPublished: boolean;
  parent: DirectoryCreate_createDirectory_directory_parent | null;
}

export interface DirectoryCreate_createDirectory {
  __typename: "DirectoryOperationResult";
  errors: DirectoryCreate_createDirectory_errors[];
  directory: DirectoryCreate_createDirectory_directory | null;
}

export interface DirectoryCreate {
  createDirectory: DirectoryCreate_createDirectory | null;
}

export interface DirectoryCreateVariables {
  name: string;
  parentId?: string | null;
}
