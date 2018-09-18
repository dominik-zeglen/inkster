import gql from "graphql-tag";

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
export interface Variables {
  input: {
    email: string;
    isActive: boolean;
  };
}
export interface Result {
  updateUser: {
    errors: Array<{
      field: string;
      message: string;
    }>;
    user: {
      id: string;
      email: string;
      updatedAt: string;
      isActive: boolean;
    };
  };
}
export default mUpdateUser;
