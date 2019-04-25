import gql from "graphql-tag";

import { TypedQuery, pageInfoFragment } from "../../api";
import { UserDetails, UserDetailsVariables } from "./types/UserDetails";

const qUser = gql`
  ${pageInfoFragment}
  query UserDetails($id: ID!, $paginate: PaginationInput!) {
    user(id: $id) {
      id
      email
      isActive
      createdAt
      updatedAt
      pages(paginate: $paginate) {
        edges {
          node {
            id
            name
            createdAt
            isPublished
          }
        }
        pageInfo {
          ...PageInfoFragment
        }
      }
    }
  }
`;
export default TypedQuery<UserDetails, UserDetailsVariables>(qUser);
