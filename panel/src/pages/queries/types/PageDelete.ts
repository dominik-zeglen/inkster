/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: PageDelete
// ====================================================

export interface PageDelete_removePage {
  __typename: "PageRemoveResult";
  removedObjectId: string;
}

export interface PageDelete {
  removePage: PageDelete_removePage | null;
}

export interface PageDeleteVariables {
  id: string;
}
