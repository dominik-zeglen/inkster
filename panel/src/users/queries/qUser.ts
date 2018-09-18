import gql from 'graphql-tag'

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
`
export interface result {
  user: {
    id: string;
    email: string;
    isActive: boolean;
    createdAt: string;
    updatedAt: string;
  }
}
export default qUser;
