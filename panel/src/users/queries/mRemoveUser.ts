import gql from "graphql-tag";

import { TypedMutation } from "../../api";
import { RemoveUser, RemoveUserVariables } from "./types/RemoveUser";

const mRemoveUser = gql`
  mutation RemoveUser($id: ID!) {
    removeUser(id: $id) {
      removedObjectId
    }
  }
`;
export default TypedMutation<RemoveUser, RemoveUserVariables>(mRemoveUser);
