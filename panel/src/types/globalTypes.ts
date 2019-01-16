/* tslint:disable */
// This file was automatically generated and should not be edited.

//==============================================================
// START Enums and Input Objects
//==============================================================

export interface PageFieldCreateInput {
  name: string;
  type: string;
  value: string;
}

export interface PageFieldUpdate {
  id: string;
  input: PageFieldUpdateInput;
}

export interface PageFieldUpdateInput {
  name?: string | null;
  value?: string | null;
}

export interface PageUpdateInput {
  name?: string | null;
  slug?: string | null;
  parentId?: string | null;
  isPublished?: boolean | null;
}

export interface PaginationInput {
  after?: string | null;
  before?: string | null;
  first?: number | null;
  last?: number | null;
}

export interface UserCreateInput {
  email: string;
  password?: string | null;
}

export interface UserUpdateInput {
  isActive?: boolean | null;
  email?: string | null;
}

export interface WebsiteUpdateInput {
  name?: string | null;
  description?: string | null;
  domain?: string | null;
}

//==============================================================
// END Enums and Input Objects
//==============================================================
