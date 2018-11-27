import gql from "graphql-tag";

import { TypedMutation } from "../../api";
import {
  DirectoryDelete,
  DirectoryDeleteVariables,
} from "./types/DirectoryDelete";

const mDirectoryDelete = gql`
  mutation DirectoryDelete($id: ID!) {
    removeDirectory(id: $id)
  }
`;
export default TypedMutation<DirectoryDelete, DirectoryDeleteVariables>(
  mDirectoryDelete,
);
