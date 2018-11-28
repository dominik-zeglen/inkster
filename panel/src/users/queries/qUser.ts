import gql from "graphql-tag";

import { TypedQuery } from "../../api";
import { UserDetails, UserDetailsVariables } from "./types/UserDetails";

const qUser = gql`
  query UserDetails($id: ID!) {
    user(id: $id) {
      id
      email
      isActive
      createdAt
      updatedAt
    }
  }
`;
export default TypedQuery<UserDetails, UserDetailsVariables>(qUser);
