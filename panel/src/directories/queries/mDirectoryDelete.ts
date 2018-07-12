import gql from "graphql-tag";

const mDirectoryDelete = gql`
  mutation DirectoryDelete($id: ID!) {
    removeContainer(id: $id)
  }
`;
export interface variables {
  id: string;
}
export default mDirectoryDelete;
