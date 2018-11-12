import gql from "graphql-tag";

const mDirectoryUpdate = gql`
  mutation DirectoryUpdate($id: ID!, $name: String!, $isPublished: Boolean) {
    updateDirectory(
      id: $id
      input: { name: $name, isPublished: $isPublished }
    ) {
      errors {
        code
        field
      }
      directory {
        id
        updatedAt
        name
        isPublished
      }
    }
  }
`;
export interface variables {
  id: string;
  name: string;
}
export default mDirectoryUpdate;
