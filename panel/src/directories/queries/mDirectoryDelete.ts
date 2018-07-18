import gql from "graphql-tag";

const mDirectoryDelete = gql`
  mutation DirectoryDelete($id: ID!) {
    removeDirectory(id: $id)
  }
`;
export interface variables {
  id: string;
}
export default mDirectoryDelete;
