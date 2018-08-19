import  gql from "graphql-tag";

const qPage = gql`
  query Page($id: ID!) {
    page(id: $id) {
      id
      createdAt
      updatedAt
      name
      slug
      fields {
        name
        type
        value
      }
      parent {
        id
      }
    }
  }
`;
export default qPage;
