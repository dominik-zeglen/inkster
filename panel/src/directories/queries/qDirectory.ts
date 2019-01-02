import gql from "graphql-tag";

import { TypedQuery, pageInfoFragment } from "../../api";
import { Directory, DirectoryVariables } from "./types/Directory";

const qDirectory = gql`
  ${pageInfoFragment}
  query Directory($id: ID!, $paginate: PaginationInput!) {
    getDirectory(id: $id) {
      id
      createdAt
      updatedAt
      name
      isPublished
      parent {
        id
      }
      pages(paginate: $paginate) {
        edges {
          node {
            id
            name
          }
        }
        pageInfo {
          ...PageInfoFragment
        }
      }
    }
  }
`;
export default TypedQuery<Directory, DirectoryVariables>(qDirectory);
