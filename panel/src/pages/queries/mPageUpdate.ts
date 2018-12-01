import gql from "graphql-tag";

import { TypedMutation } from "../../api";
import { PageUpdate, PageUpdateVariables } from "./types/PageUpdate";

const mPageUpdate = gql`
  mutation PageUpdate(
    $id: ID!
    $input: PageUpdateInput
    $add: [PageFieldCreateInput!]
    $update: [PageFieldUpdate!]
    $remove: [String!]
  ) {
    updatePage(
      id: $id
      input: $input
      addFields: $add
      updateFields: $update
      removeFields: $remove
    ) {
      errors {
        code
        field
      }
      page {
        id
        updatedAt
        name
        slug
        isPublished
        fields {
          id
          name
          type
          value
        }
      }
    }
  }
`;
export default TypedMutation<PageUpdate, PageUpdateVariables>(mPageUpdate);
