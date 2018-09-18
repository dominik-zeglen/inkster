import gql from "graphql-tag";

const mRemoveUser = gql`
  mutation RemoveUser($id: ID!) {
    removeUser(id: $id) {
      removedObjectId
    }
  }
`;
export interface Variables {
  id: string;
}
export interface Result {
  removeUser: {
    removedObjectId: string;
  };
}
export default mRemoveUser;
