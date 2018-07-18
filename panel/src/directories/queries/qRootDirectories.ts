import  gql from "graphql-tag";

const qRootDirectories = gql`
  query RootDirectories {
    getRootDirectories {
      id
      name
    }
  }
`;
export default qRootDirectories;
