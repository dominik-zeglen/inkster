import gql from "graphql-tag";

import { TypedQuery } from "../../api";
import { UserList } from "./types/UserList";

const qUsers = gql`
  query UserList {
    users {
      id
      email
      isActive
    }
  }
`;
export default TypedQuery<UserList, {}>(qUsers);
