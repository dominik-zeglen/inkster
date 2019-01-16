import gql from "graphql-tag";

import { TypedMutation } from "../../api";
import { WebsiteUpdate, WebsiteUpdateVariables } from "./types/WebsiteUpdate";

const mWebsiteUpdate = gql`
  mutation WebsiteUpdate($input: WebsiteUpdateInput!) {
    updateWebsite(input: $input) {
      errors {
        code
        field
      }
      website {
        name
        description
        domain
      }
    }
  }
`;
export default TypedMutation<WebsiteUpdate, WebsiteUpdateVariables>(
  mWebsiteUpdate,
);
