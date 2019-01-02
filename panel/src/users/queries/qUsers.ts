import gql from "graphql-tag";

import { TypedQuery, pageInfoFragment } from "../../api";
import { UserList, UserListVariables } from "./types/UserList";

const qUsers = gql`
  ${pageInfoFragment}
  query UserList($paginate: PaginationInput!) {
    users(paginate: $paginate) {
      edges {
        node {
          id
          email
          isActive
        }
      }
      pageInfo {
        ...PageInfoFragment
      }
    }
  }
`;
export default TypedQuery<UserList, UserListVariables>(qUsers);
