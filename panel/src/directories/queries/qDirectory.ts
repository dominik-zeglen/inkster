import  gql from "graphql-tag";

const qDirectory = gql`
  query Directory($id: ID!) {
    getDirectory(id: $id) {
      id
      name
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
