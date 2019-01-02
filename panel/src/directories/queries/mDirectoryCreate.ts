import gql from "graphql-tag";

import {
  DirectoryCreate,
  DirectoryCreateVariables,
} from "./types/DirectoryCreate";
import { TypedMutation } from "../../api";

const mDirectoryCreate = gql`
  mutation DirectoryCreate($name: String!, $parentId: ID) {
    createDirectory(input: { name: $name, parentId: $parentId }) {
      errors {
        field
        code
      }
      directory {
        id
        createdAt
        updatedAt
        name
        isPublished
        parent {
          id
        }
      }
    }
  }
`;
export default TypedMutation<DirectoryCreate, DirectoryCreateVariables>(
  mDirectoryCreate,
);
