import gql from "graphql-tag";

const mPageCreate = gql`
  mutation PageCreate(
    $parentId: ID!
    $name: String!
    $fields: [PageFieldCreateInput!]
  ) {
    createPage(input: { name: $name, parentId: $parentId, fields: $fields }) {
      userErrors {
        field
        message
      }
      page {
        id
        name
        slug
        fields {
          name
          type
          value
        }
      }
    }
  }
`;
export interface variables {
  name: string;
  parentId: string;
  fields?: Array<{
    name: string;
    type: string;
    value: string;
  }>;
}
export interface result {
  userErrors: Array<{
    field: string;
    message: string;
  }>;
  page: {
    id: string;
    name: string;
    slug: string;
    fields: Array<{
      name: string;
      type: string;
      value: string;
    }>;
  };
}
export default mPageCreate;
