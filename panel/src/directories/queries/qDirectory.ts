import  gql from "graphql-tag";

const qDirectory = gql`
  query Directory($id: ID!) {
    getDirectory(id: $id) {
      id
      createdAt
      updatedAt
      name
      isPublished
      parent {
        id
      }
      pages {
        id
        name
      }
    }
  }
`;
export default qDirectory;
