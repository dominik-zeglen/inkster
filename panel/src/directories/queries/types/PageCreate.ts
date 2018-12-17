/* tslint:disable */
// This file was automatically generated and should not be edited.

import { PageFieldCreateInput } from "./../../../types/globalTypes";

// ====================================================
// GraphQL mutation operation: PageCreate
// ====================================================

export interface PageCreate_createPage_errors {
  __typename: "InputError";
  field: string;
  code: number;
}

export interface PageCreate_createPage_page_fields {
  __typename: "PageField";
  name: string;
  type: string;
  value: string | null;
}

export interface PageCreate_createPage_page {
  __typename: "Page";
  id: string;
  createdAt: string;
  updatedAt: string;
  name: string;
  slug: string;
  isPublished: boolean;
  fields: (PageCreate_createPage_page_fields | null)[] | null;
}

export interface PageCreate_createPage {
  __typename: "PageOperationResult";
  errors: PageCreate_createPage_errors[];
  page: PageCreate_createPage_page | null;
}

export interface PageCreate {
  createPage: PageCreate_createPage | null;
}

export interface PageCreateVariables {
  parentId: string;
  name: string;
  fields?: PageFieldCreateInput[] | null;
}
