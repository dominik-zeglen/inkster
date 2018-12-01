import gql from "graphql-tag";

import { TypedQuery } from "../../api";
import { RootDirectories } from "./types/RootDirectories";

const qRootDirectories = gql`
  query RootDirectories {
    getRootDirectories {
      id
      name
      isPublished
    }
  }
`;
export default TypedQuery<RootDirectories, {}>(qRootDirectories);
