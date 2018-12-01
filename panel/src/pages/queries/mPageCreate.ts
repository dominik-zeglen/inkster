import gql from "graphql-tag";

import { TypedMutation } from "../../api";
import { PageCreate, PageCreateVariables } from "./types/PageCreate";

const mPageCreate = gql`
  mutation PageCreate(
    $parentId: ID!
    $name: String!
    $fields: [PageFieldCreateInput!]
  ) {
    createPage(input: { name: $name, parentId: $parentId, fields: $fields }) {
      errors {
        field
        code
      }
      page {
        id
        createdAt
        updatedAt
        name
        slug
        isPublished
        fields {
          name
          type
          value
        }
      }
    }
  }
`;
export default TypedMutation<PageCreate, PageCreateVariables>(mPageCreate);
