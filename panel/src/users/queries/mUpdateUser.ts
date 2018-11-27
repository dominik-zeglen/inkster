import gql from "graphql-tag";
import { TypedMutation } from "../../api";
import { UpdateUser, UpdateUserVariables } from "./types/UpdateUser";

const mUpdateUser = gql`
  mutation UpdateUser($id: ID!, $input: UserUpdateInput!) {
    updateUser(id: $id, input: $input) {
      errors {
        field
        message
      }
      user {
        id
        email
        updatedAt
        isActive
      }
    }
  }
`;
export default TypedMutation<UpdateUser, UpdateUserVariables>(mUpdateUser);
