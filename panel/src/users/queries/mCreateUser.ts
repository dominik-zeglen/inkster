import gql from "graphql-tag";

import { TypedMutation } from "../../api";
import { CreateUser, CreateUserVariables } from "./types/CreateUser";

const mCreateUser = gql`
  mutation CreateUser($input: UserCreateInput!) {
    createUser(input: $input, sendInvitation: true) {
      errors {
        field
        message
      }
      user {
        id
        email
        createdAt
        updatedAt
        isActive
      }
    }
  }
`;
export default TypedMutation<CreateUser, CreateUserVariables>(mCreateUser);
