import  gql from "graphql-tag";

const qRootDirectories = gql`
  query RootDirectories {
    getRootContainers {
      id
      name
    }
  }
`;
export default qRootDirectories;
