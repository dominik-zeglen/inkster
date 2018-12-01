import gql from "graphql-tag";

import { TypedQuery } from "../../api";
import { Directory, DirectoryVariables } from "./types/Directory";

const qDirectory = gql`
  query Directory($id: ID!) {
    getDirectory(id: $id) {
      id
      createdAt
      updatedAt
      name
      isPublished
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
export default TypedQuery<Directory, DirectoryVariables>(qDirectory);
