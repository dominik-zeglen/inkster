/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: Page
// ====================================================

export interface Page_page_fields {
  __typename: "PageField";
  id: string;
  name: string;
  type: string;
  value: string | null;
}

export interface Page_page_parent {
  __typename: "Directory";
  id: string;
}

export interface Page_page {
  __typename: "Page";
  id: string;
  createdAt: string;
  updatedAt: string;
  name: string;
  slug: string;
  isPublished: boolean;
  fields: (Page_page_fields | null)[] | null;
  parent: Page_page_parent;
}

export interface Page {
  page: Page_page | null;
}

export interface PageVariables {
  id: string;
}
