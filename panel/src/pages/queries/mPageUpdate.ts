import gql from "graphql-tag";

const mPageUpdate = gql`
  mutation PageUpdate(
    $id: ID!
    $input: PageUpdateInput
    $add: [PageFieldCreateInput!]
    $remove: [String!]
  ) {
    updatePage(id: $id, input: $input, addFields: $add, removeFields: $remove) {
      userErrors {
        field
        message
      }
      page {
        id
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
  id: string;
  input: {
    name: string;
    slug: string;
    isPublished?: boolean;
    fields: Array<{
      name: string;
      update: {
        name: string;
        value: string;
      };
    }>;
  };
  add: Array<{
    name: string;
    type: string;
    value: string;
  }>;
  remove: string[];
}
export default mPageUpdate;
