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

export interface PageUpdateInput {
  name?: string | null;
  slug?: string | null;
  parentId?: string | null;
  isPublished?: boolean | null;
}

export interface UserCreateInput {
  email: string;
  password?: string | null;
}

export interface UserUpdateInput {
  isActive?: boolean | null;
  email?: string | null;
}

//==============================================================
// END Enums and Input Objects
//==============================================================
