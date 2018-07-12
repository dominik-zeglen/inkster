import gql from "graphql-tag";

const mDirectoryUpdate = gql`
  mutation PageUpdate($id: ID!, $name: String) {
    updatePage(id: $id, input: { name: $name }) {

      id
      name
    }
  }
`;
export interface variables {
  id: string;
  name: string;
}
export default mDirectoryUpdate;
