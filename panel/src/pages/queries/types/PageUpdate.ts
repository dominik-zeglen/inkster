/* tslint:disable */
// This file was automatically generated and should not be edited.

import { PageUpdateInput, PageFieldCreateInput, PageFieldUpdate } from "./../../../types/globalTypes";

// ====================================================
// GraphQL mutation operation: PageUpdate
// ====================================================

export interface PageUpdate_updatePage_errors {
  __typename: "InputError";
  code: number;
  field: string;
}

export interface PageUpdate_updatePage_page_fields {
  __typename: "PageField";
  id: string;
  name: string;
  type: string;
  value: string | null;
}

export interface PageUpdate_updatePage_page {
  __typename: "Page";
  id: string;
  updatedAt: string;
  name: string;
  slug: string;
  isPublished: boolean;
  fields: (PageUpdate_updatePage_page_fields | null)[] | null;
}

export interface PageUpdate_updatePage {
  __typename: "PageOperationResult";
  errors: PageUpdate_updatePage_errors[];
  page: PageUpdate_updatePage_page | null;
}

export interface PageUpdate {
  updatePage: PageUpdate_updatePage | null;
}

export interface PageUpdateVariables {
  id: string;
  input?: PageUpdateInput | null;
  add?: PageFieldCreateInput[] | null;
  update?: PageFieldUpdate[] | null;
  remove?: string[] | null;
}
