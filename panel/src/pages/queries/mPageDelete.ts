import gql from "graphql-tag";

const mPageDelete = gql`
  mutation PageDelete($id: ID!) {
    removePage(id: $id)
  }
`;
export interface variables {
  id: string;
}
export default mPageDelete;
