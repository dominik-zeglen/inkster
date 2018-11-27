import gql from "graphql-tag";
import { TypedMutation } from "../../api";
import {
  DirectoryUpdate,
  DirectoryUpdateVariables,
} from "./types/DirectoryUpdate";

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
export default TypedMutation<DirectoryUpdate, DirectoryUpdateVariables>(
  mDirectoryUpdate,
);
