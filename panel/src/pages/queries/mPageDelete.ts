import gql from "graphql-tag";

import { TypedMutation } from "../../api";
import { PageDelete, PageDeleteVariables } from "./types/PageDelete";

const mPageDelete = gql`
  mutation PageDelete($id: ID!) {
    removePage(id: $id) {
      removedObjectId
    }
  }
`;
export default TypedMutation<PageDelete, PageDeleteVariables>(mPageDelete);
