import gql from "graphql-tag";

import { TypedQuery, pageInfoFragment } from "../../api";
import {
  RootDirectories,
  RootDirectoriesVariables,
} from "./types/RootDirectories";

const qRootDirectories = gql`
  ${pageInfoFragment}
  query RootDirectories($paginate: PaginationInput!) {
    getRootDirectories(paginate: $paginate) {
      edges {
        node {
          id
          name
          isPublished
        }
      }
      pageInfo {
        ...PageInfoFragment
      }
    }
  }
`;
export default TypedQuery<RootDirectories, RootDirectoriesVariables>(
  qRootDirectories,
);
