import  gql from "graphql-tag";

const qRootDirectories = gql`
  query RootDirectories {
    getRootDirectories {
      id
      name
      isPublished
    }
  }
`;
export default qRootDirectories;
