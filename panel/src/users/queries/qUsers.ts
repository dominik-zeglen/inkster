import gql from 'graphql-tag'

const qUsers = gql`
  query UserList {
    users {
      id
      email
      isActive
    }
  }
`
export interface result {
  users: Array<{
    id: string;
    email: string;
    isActive: boolean;
  }>
}
export default qUsers;
