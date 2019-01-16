/* tslint:disable */
// This file was automatically generated and should not be edited.

import { WebsiteUpdateInput } from "./../../../types/globalTypes";

// ====================================================
// GraphQL mutation operation: WebsiteUpdate
// ====================================================

export interface WebsiteUpdate_updateWebsite_errors {
  __typename: "InputError";
  code: number;
  field: string;
}

export interface WebsiteUpdate_updateWebsite_website {
  __typename: "Website";
  name: string;
  description: string;
  domain: string;
}

export interface WebsiteUpdate_updateWebsite {
  __typename: "WebsiteOperationResult";
  errors: WebsiteUpdate_updateWebsite_errors[];
  website: WebsiteUpdate_updateWebsite_website | null;
}

export interface WebsiteUpdate {
  updateWebsite: WebsiteUpdate_updateWebsite;
}

export interface WebsiteUpdateVariables {
  input: WebsiteUpdateInput;
}
