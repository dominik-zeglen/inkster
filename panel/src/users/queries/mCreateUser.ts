import gql from 'graphql-tag'

const mCreateUser = gql`
  mutation CreateUser($input: UserCreateInput!) {
    createUser(input: $input) {
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
`
export interface variables {
  input: {
    email: string;
  }
}
export interface result {
  errors: Array<{
    field: string;
    message: string;
  }>
  user: {
    id: string;
    email: string;
    createdAt: string;
    updatedAt: string;
    isActive: boolean;
  }
}
export default mCreateUser;
