import gql from "graphql-tag";

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
  errors: Array<{
    field: string;
    code: number;
  }>;
  page: {
    id: string;
    name: string;
    slug: string;
    isPublished: boolean;
    fields: Array<{
      name: string;
      type: string;
      value: string;
    }>;
  };
}
export default mPageCreate;
